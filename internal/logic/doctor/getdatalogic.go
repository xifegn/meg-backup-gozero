package doctor

import (
	"context"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDataLogic {
	return &GetDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDataLogic) GetData(req *types.GetDataRequest) (resp *types.GetDataResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
