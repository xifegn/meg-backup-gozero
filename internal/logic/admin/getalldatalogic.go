package admin

import (
	"context"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllDataLogic {
	return &GetAllDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllDataLogic) GetAllData() (resp *types.GetAllDataResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
