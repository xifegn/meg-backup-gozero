package admin

import (
	"context"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminGetAllDoctorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminGetAllDoctorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminGetAllDoctorLogic {
	return &AdminGetAllDoctorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminGetAllDoctorLogic) AdminGetAllDoctor() (resp *types.AdminGetAllDoctorResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
