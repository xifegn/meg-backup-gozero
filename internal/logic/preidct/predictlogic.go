package preidct

import (
	"context"
	"fmt"
	"github.com/levigross/grequests"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PredictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type PredictInfo struct {
	Message string `json:"message"`
}

func NewPredictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PredictLogic {
	return &PredictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PredictLogic) Predict(req *types.PredictRequest) (resp *types.PredictResponse, err error) {
	url := fmt.Sprint(l.svcCtx.Config.Url + "/predict")
	response, err := grequests.Get(url, &grequests.RequestOptions{
		Params:  map[string]string{"code": req.Code},
		Context: l.ctx,
	})
	if err != nil {
		return nil, err
	}
	if response.Ok {
		var data PredictInfo
		err := response.JSON(&data)
		if err != nil {
			return nil, err
		}
		return &types.PredictResponse{Message: data.Message}, nil
	}
	return &types.PredictResponse{}, nil
}
