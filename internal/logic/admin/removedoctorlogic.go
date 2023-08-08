package admin

import (
	"context"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveDoctorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveDoctorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveDoctorLogic {
	return &RemoveDoctorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveDoctorLogic) RemoveDoctor(req *types.RemoveDoctorRequest) (resp *types.RemoveDoctorResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
