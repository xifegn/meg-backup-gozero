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

func (l *AdminGetAllDoctorLogic) AdminGetAllDoctor() (resp []*types.AdminGetAllDoctorResponse, err error) {
	res, err := l.svcCtx.DoctorModel.AdminGetAllDoctor(l.ctx, 0)
	if err != nil {
		return nil, err
	}
	doctorInfo := make([]*types.AdminGetAllDoctorResponse, 0)
	for _, item := range res {
		doctorInfo = append(doctorInfo, &types.AdminGetAllDoctorResponse{
			Id:           item.Id,
			Name:         item.Name,
			RegisterTime: item.RegisterTime,
			Telephone:    item.Telephone,
			Username:     item.Username,
		})
	}
	return doctorInfo, nil
}
