package admin

import (
	"context"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuotaApplyInsertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuotaApplyInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuotaApplyInsertLogic {
	return &QuotaApplyInsertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuotaApplyInsertLogic) QuotaApplyInsert(req *types.QuotaApplyInsertRequest) (resp *types.QuotaApplyInsertResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
