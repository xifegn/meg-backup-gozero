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
	"time"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.ChangePasswordRequest) (resp *types.ChangePasswordResponse, err error) {
	res, err := l.svcCtx.DoctorModel.GetAllDoctorByUsername(l.ctx, req.Username)
	if err != nil {
		return nil, err
	}
	fmt.Println(res[0])
	//if req.OldPassword != res[0].Password {
	//	return nil, errors.New("old password error")
	//}
	if cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.OldPassword) != res[0].Password {
		return nil, errors.New("old password error")
	}
	//input := res[0].UploadTime[:10]
	//t, err := time.Parse("2012-01-02", input)
	//if err != nil {
	//	return nil, err
	//}
	newPassword := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.NewPassword)
	updateDoctor := doctor.Doctor{
		Id:          res[0].Id,
		Username:    res[0].Username,
		Password:    newPassword,
		IsAdmin:     res[0].IsAdmin,
		Name:        res[0].Name,
		Phonenumber: res[0].Phonenumber,
		UploadTime:  time.Now(),
	}
	err = l.svcCtx.DoctorModel.Update(l.ctx, &updateDoctor)
	if err != nil {
		return nil, err
	}
	return &types.ChangePasswordResponse{
		Message: "Ok",
	}, nil
}
