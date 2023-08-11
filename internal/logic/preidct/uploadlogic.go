package preidct

import (
	"context"
	"fmt"
	"net/http"

	"github.com/levigross/grequests"
	"github.com/zeromicro/go-zero/core/logx"
	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"
)

const maxFileSize = 100 << 20

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

type TestInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewUploadLogic(r *http.Request, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(r.Context()),
		r:      r,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req *types.UploadRequest) (resp *types.UploadResponse, err error) {
	l.r.ParseMultipartForm(maxFileSize)
	files := l.r.MultipartForm.File["myFile"]
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
			FieldName:    "myFile",
		}
		fileUploads = append(fileUploads, fileUpload)
	}
	url := fmt.Sprint(l.svcCtx.Config.Url + "/upload")
	response, err := grequests.Post(url, &grequests.RequestOptions{
		Params:  map[string]string{"code": req.Code},
		Files:   fileUploads,
		Context: l.ctx,
	})
	if err != nil {
		return nil, err
	}
	defer response.Close()
	if response.Ok {
		var data TestInfo
		err := response.JSON(&data)
		if err != nil {
			return nil, err
		}
		return &types.UploadResponse{
			Base: types.Base{Code: 0, Msg: "OK"},
			Data: types.Message{Message: data.Message},
		}, nil
	}
	return &types.UploadResponse{}, nil
}
