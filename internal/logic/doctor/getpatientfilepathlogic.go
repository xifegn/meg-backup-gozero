package doctor

import (
	"context"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPatientFilePathLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPatientFilePathLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPatientFilePathLogic {
	return &GetPatientFilePathLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPatientFilePathLogic) GetPatientFilePath(req *types.GetPatientFilePathRequest) (resp *types.GetPatientFilePathResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
