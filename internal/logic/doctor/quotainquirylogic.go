package doctor

import (
	"context"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuotaInquiryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuotaInquiryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuotaInquiryLogic {
	return &QuotaInquiryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuotaInquiryLogic) QuotaInquiry(req *types.QuotaInquiryRequest) (resp *types.QuotaInquiryResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
