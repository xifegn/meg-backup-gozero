package doctor

import (
	"context"
	"meg-backup-gozero/models/doctor"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertDoctorInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertDoctorInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertDoctorInfoLogic {
	return &InsertDoctorInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertDoctorInfoLogic) InsertDoctorInfo(req *types.InsertDoctorInfoRequest) (resp *types.InsertDoctorInfoResponse, err error) {
	res, _ := l.svcCtx.DoctorModel.HaveDoctorInfo(l.ctx, req.Did)
	if res == -100 {
		newInfo := doctor.DoctorInfo{
			Email:           req.Email,
			Nickname:        req.Nickname,
			Regions:         req.Regions,
			SecretKey:       l.svcCtx.Config.Salt,
			SelfInformation: req.SelfInformation,
			Did:             req.Did,
		}
		_, err = l.svcCtx.DoctorInfoModel.Insert(l.ctx, &newInfo)
		if err != nil {
			return nil, err
		}
		return &types.InsertDoctorInfoResponse{Message: "Insert success"}, nil
	}
	updateInfo := doctor.DoctorInfo{
		Id:              res,
		Email:           req.Email,
		Nickname:        req.Nickname,
		Regions:         req.Regions,
		SelfInformation: req.SelfInformation,
		SecretKey:       l.svcCtx.Config.Salt,
		Did:             req.Did,
	}
	err = l.svcCtx.DoctorInfoModel.Update(l.ctx, &updateInfo)
	if err != nil {
		return nil, err
	}

	return &types.InsertDoctorInfoResponse{Message: "update success"}, nil
}
