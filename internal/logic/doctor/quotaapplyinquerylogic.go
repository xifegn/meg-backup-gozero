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

func (l *QuotaApplyInqueryLogic) QuotaApplyInquery() (resp []*types.QuotaApplyInqueryResponse, err error) {
	res, _ := l.svcCtx.QuotaApplyModel.GetQuotaApply(l.ctx)
	data := make([]*types.QuotaApplyInqueryResponse, 0)
	for _, item := range res {
		data = append(data, &types.QuotaApplyInqueryResponse{
			Id:          item.Id,
			Username:    item.Username,
			QuotaAmount: item.QuotaAmount,
			Amount:      item.Amount,
			CreatedAt:   item.CreatedAt,
		})
	}
	return data, nil
}
