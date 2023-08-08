package preidct

import (
	"context"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PredictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPredictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PredictLogic {
	return &PredictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PredictLogic) Predict(req *types.PredictRequest) (resp *types.PredictResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
