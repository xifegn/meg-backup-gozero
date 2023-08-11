package doctor

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"meg-backup-gozero/common/cryptx"
	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"
	"meg-backup-gozero/models/doctor"
)

type ChangeEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeEmailLogic {
	return &ChangeEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeEmailLogic) ChangeEmail(req *types.ChangeEmailRequest) (resp *types.ChangeEmailResponse, err error) {
	res, err := l.svcCtx.DoctorModel.GetAllDoctorByUsername(l.ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password) != res[0].Password {
		return nil, errors.New("password error")
	}
	infoRes, _ := l.svcCtx.DoctorInfoModel.GetAllDoctorInfoByUsername(l.ctx, req.Username)
	if err != nil {
		return nil, errors.New("user not found")
	}
	fmt.Println(infoRes[0])
	updateInfo := doctor.DoctorInfo{
		Id:              infoRes[0].Id,
		Email:           req.Email,
		Nickname:        infoRes[0].NickName,
		Regions:         infoRes[0].Regions,
		SelfInformation: infoRes[0].SelfInformation,
		SecretKey:       infoRes[0].SecretKey,
		Did:             infoRes[0].Did,
	}
	err = l.svcCtx.DoctorInfoModel.Update(l.ctx, &updateInfo)
	if err != nil {
		return nil, err
	}
	return &types.ChangeEmailResponse{Message: "update success"}, nil
}
