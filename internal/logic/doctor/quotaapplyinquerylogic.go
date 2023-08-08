package doctor

import (
	"context"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuotaApplyInqueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuotaApplyInqueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuotaApplyInqueryLogic {
	return &QuotaApplyInqueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuotaApplyInqueryLogic) QuotaApplyInquery() (resp *types.QuotaApplyInqueryResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
