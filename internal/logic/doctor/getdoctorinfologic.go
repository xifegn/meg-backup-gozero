package doctor

import (
	"context"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDoctorInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDoctorInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDoctorInfoLogic {
	return &GetDoctorInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDoctorInfoLogic) GetDoctorInfo(req *types.GetDoctorInfoRequest) (resp *types.GetDoctorInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
