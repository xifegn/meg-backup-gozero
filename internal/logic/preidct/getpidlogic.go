package preidct

import (
	"context"
	"fmt"
	"github.com/levigross/grequests"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPidLogic {
	return &GetPidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPidLogic) GetPid() (resp *types.GetPidResponse, err error) {
	url := fmt.Sprint(l.svcCtx.Config.Url + "/getpid")
	response, err := grequests.Get(url, nil)
	if err != nil {
		return nil, err
	}
	defer response.Close()
	if response.Ok {
		var data types.GetPidResponse
		err := response.JSON(&data)
		//fmt.Println(data)
		if err != nil {
			return nil, err
		}
		return &types.GetPidResponse{
			Code:       200,
			Msg:        data.Msg,
			CurrentPid: data.CurrentPid,
		}, nil
	}

	return &types.GetPidResponse{}, nil
}
