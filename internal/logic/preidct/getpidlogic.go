package preidct

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
