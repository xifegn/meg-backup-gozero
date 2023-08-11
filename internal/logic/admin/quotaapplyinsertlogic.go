package admin

import (
	"context"
	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"
	"meg-backup-gozero/models/quota"
	"time"

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
	newData := quota.QuotaApply{
		Amount:    req.Amount,
		Username:  req.Username,
		CreatedAt: time.Now(),
	}
	_, err = l.svcCtx.QuotaApplyModel.Insert(l.ctx, &newData)
	if err != nil {
		return nil, err
	}
	return &types.QuotaApplyInsertResponse{Message: "Insert success"}, nil
}
