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
	res, err := l.svcCtx.DoctorInfoModel.DoctorInfoByUsername(l.ctx, req.Username)
	if err != nil {
		return nil, err
	}

	return &types.GetDoctorInfoResponse{
		Username:        res[0].Username,
		Phonenumber:     res[0].Phonenumber,
		Id:              res[0].Id,
		SecretKey:       res[0].SecretKey,
		Email:           res[0].Email,
		Nickname:        res[0].Nickname,
		Regions:         res[0].Regions,
		SelfInformation: res[0].SelfInformation,
	}, nil
}
