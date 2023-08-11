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
	err = l.svcCtx.PatientModel.Delete(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &types.RemovePatientResponse{Message: "OK"}, nil
}
