package preidct

import (
	"context"
	"fmt"
	"github.com/levigross/grequests"
	"net/http"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SmUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}
type dataResp struct {
	Url string `json:"url"`
}

func NewSmUploadLogic(r *http.Request, svcCtx *svc.ServiceContext) *SmUploadLogic {
	return &SmUploadLogic{
		Logger: logx.WithContext(r.Context()),
		r:      r,
		svcCtx: svcCtx,
	}
}

func (l *SmUploadLogic) SmUpload() (resp *types.SmUploadResponse, err error) {
	l.r.ParseMultipartForm(maxFileSize)
	files := l.r.MultipartForm.File["smfile"]
	var fileUploads []grequests.FileUpload
	// 用一个for循环来遍历files切片
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		logx.Infof("upload file: %+v, file size: %d, MIME header: %+v",
			fileHeader.Filename, fileHeader.Size, fileHeader.Header)

		fileUpload := grequests.FileUpload{
			FileName:     fileHeader.Filename,
			FileContents: file,
			FileMime:     fileHeader.Header.Get("Content-Type"),
			FieldName:    "smfile",
		}
		fileUploads = append(fileUploads, fileUpload)
	}
	url := fmt.Sprint(l.svcCtx.Config.Url + "/sm_upload")
	response, err := grequests.Post(url, &grequests.RequestOptions{
		Files:   fileUploads,
		Context: l.ctx,
	})
	if err != nil {
		return nil, err
	}
	defer response.Close()
	if response.Ok {
		var data dataResp
		err := response.JSON(&data)
		if err != nil {
			return nil, err
		}
		return &types.SmUploadResponse{
			Base: types.Base{
				Code: 0,
				Msg:  "OK",
			},
			Data: types.Url{
				Url: data.Url,
			},
		}, nil
	}
	return &types.SmUploadResponse{}, nil
}
