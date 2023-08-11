package preidct

import (
	"context"
	"fmt"
	"github.com/levigross/grequests"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendPidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendPidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendPidLogic {
	return &SendPidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendPidLogic) SendPid(req *types.SendPidRequest) (resp *types.SendPidResponse, err error) {
	url := fmt.Sprint(l.svcCtx.Config.Url + "/sendpid")
	response, err := grequests.Post(url, &grequests.RequestOptions{
		JSON:    types.SendPidRequest{req.CurrentPid},
		Context: l.ctx,
	})
	defer response.Close()
	if err != nil {
		return nil, err
	}
	if response.Ok {
		var data types.SendPidResponse
		err := response.JSON(&data)
		if err != nil {
			return nil, err
		}
		return &types.SendPidResponse{
			Code: 200,
			Msg:  data.Msg,
		}, nil
	}

	return &types.SendPidResponse{}, nil
}
