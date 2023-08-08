package doctor

import (
	"context"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuotaApplyRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuotaApplyRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuotaApplyRemoveLogic {
	return &QuotaApplyRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuotaApplyRemoveLogic) QuotaApplyRemove(req *types.QuotaApplyRemoveRequest) (resp *types.QuotaApplyRemoveResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
