package admin

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"
	"meg-backup-gozero/models/quota"
	"time"
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
	res, _ := l.svcCtx.QuotaModel.AddQuotaByUsername(l.ctx, req.Username)
	fmt.Println(res)
	if res == nil {
		quotaData := quota.Quota{
			Amount:    req.Amount + 100,
			CreatedAt: time.Now(),
			Username:  req.Username,
		}
		_, err := l.svcCtx.QuotaModel.Insert(l.ctx, &quotaData)
		if err != nil {
			return nil, err
		}
		return &types.AddQuotaResponse{
			Message: "Insert Ok",
		}, nil
	}
	updateQuota := quota.Quota{
		Id:        res[0].Id,
		Amount:    req.Amount + res[0].Amount,
		CreatedAt: time.Now(),
		Username:  req.Username,
	}
	err = l.svcCtx.QuotaModel.Update(l.ctx, &updateQuota)
	if err != nil {
		return nil, err
	}
	return &types.AddQuotaResponse{Message: "Add quota success"}, nil
}
