package preidct

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
