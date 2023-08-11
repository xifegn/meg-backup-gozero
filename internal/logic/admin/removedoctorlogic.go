package admin

import (
	"context"
	"errors"
	"fmt"
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
	res, err := l.svcCtx.DoctorModel.FindIdByUsername(l.ctx, req.Username)
	fmt.Println(res)
	if err != nil {
		return nil, errors.New("user not found")
	}
	fmt.Println(res)
	err = l.svcCtx.DoctorModel.Delete(l.ctx, res)
	if err != nil {
		return nil, err
	}
	return &types.RemoveDoctorResponse{Message: "OK"}, nil
}
