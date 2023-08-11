package doctor

import (
	"context"
	"fmt"
	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"
	"strconv"

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

func (l *QuotaInquiryLogic) QuotaInquiry(req *types.QuotaInquiryRequest) (resp []*types.QuotaInquiryResponse, err error) {
	res, err := l.svcCtx.QuotaModel.GetQuotaForAWeekByUsername(l.ctx, req.Username)
	if err != nil {
		return nil, err
	}
	data := make([]*types.QuotaInquiryResponse, 0, 8)
	weekdays := []string{"周一", "周二", "周三", "周四", "周五", "周六", "周日"}
	for i := 0; i < len(res); i++ {
		data = append(data, &types.QuotaInquiryResponse{
			DailyCallLimit: strconv.FormatInt(res[i].DailyCallLimit, 10),
			CallNumbers:    strconv.FormatInt(res[i].FileCount, 10),
			CallVolumes:    fmt.Sprintf("%.1f", float64(res[i].FileCount)/float64(res[i].DailyCallLimit)*100),
			Type:           weekdays[i],
		})

	}
	return data, nil
}
