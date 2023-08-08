package admin

import (
	"context"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InformationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInformationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InformationLogic {
	return &InformationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InformationLogic) Information() (resp *types.ChartDataResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
