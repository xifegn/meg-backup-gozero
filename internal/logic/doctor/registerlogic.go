package doctor

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"meg-backup-gozero/common/cryptx"
	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"
	"meg-backup-gozero/models/doctor"
	"time"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	_, err = l.svcCtx.DoctorModel.FindOneByUsername(l.ctx, req.Username)
	if err == nil {
		return nil, status.Error(500, "user exist")
	}

	if err == doctor.ErrNotFound {
		newDoctor := doctor.Doctor{
			Username:    req.Username,
			Password:    cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password),
			IsAdmin:     req.IsAdmin,
			Name:        req.Name,
			Phonenumber: req.Number,
			UploadTime:  time.Now(),
		}
		res, err := l.svcCtx.DoctorModel.Insert(l.ctx, &newDoctor)
		if err != nil {
			return nil, err
		}
		newDoctor.Id, err = res.LastInsertId()
		return &types.RegisterResponse{
			Code:    200,
			Message: "Ok",
		}, nil
	}
	return &types.RegisterResponse{}, nil
}
