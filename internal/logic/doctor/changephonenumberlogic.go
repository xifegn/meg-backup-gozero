package doctor

import (
	"context"
	"errors"
	"meg-backup-gozero/common/cryptx"
	"meg-backup-gozero/models/doctor"
	"time"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePhoneNumberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePhoneNumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePhoneNumberLogic {
	return &ChangePhoneNumberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePhoneNumberLogic) ChangePhoneNumber(req *types.ChangePhoneNumberRequest) (resp *types.ChangePhoneNumberResponse, err error) {
	res, err := l.svcCtx.DoctorModel.GetAllDoctorByUsername(l.ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password) != res[0].Password {
		return nil, errors.New("password error")
	}
	updateDoctor := doctor.Doctor{
		Id:          res[0].Id,
		Username:    res[0].Username,
		Password:    res[0].Password,
		IsAdmin:     res[0].IsAdmin,
		Name:        res[0].Name,
		Phonenumber: req.PhoneNumber,
		UploadTime:  time.Now(),
	}
	err = l.svcCtx.DoctorModel.Update(l.ctx, &updateDoctor)
	if err != nil {
		return nil, err
	}
	return &types.ChangePhoneNumberResponse{Message: "OK"}, nil
}
