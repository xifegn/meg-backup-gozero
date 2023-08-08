package admin

import (
	"context"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemovePatientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemovePatientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemovePatientLogic {
	return &RemovePatientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemovePatientLogic) RemovePatient(req *types.RemovePatientRequest) (resp *types.RemovePatientResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
