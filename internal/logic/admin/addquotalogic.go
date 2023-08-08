package admin

import (
	"context"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddQuotaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddQuotaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddQuotaLogic {
	return &AddQuotaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddQuotaLogic) AddQuota(req *types.AddQuotaRequest) (resp *types.AddQuotaResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
