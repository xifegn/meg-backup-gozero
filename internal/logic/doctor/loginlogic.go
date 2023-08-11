package doctor

import (
	"context"
	"google.golang.org/grpc/status"
	"meg-backup-gozero/common/cryptx"
	"meg-backup-gozero/common/jwtx"
	"meg-backup-gozero/models/doctor"
	"time"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	res, err := l.svcCtx.DoctorModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		if err == doctor.ErrNotFound {
			return nil, status.Error(100, "user not found")
		}
		return nil, status.Error(500, err.Error())
	}
	if res.Password != cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password) {
		return nil, status.Error(100, "password error")
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, res.Id)
	if err != nil {
		return nil, err
	}

	var token string
	switch res.IsAdmin {
	case 1:
		token = "admin"
	case 0:
		token = "editor"
	}

	return &types.LoginResponse{
		AccessToken:  accessToken,
		AccessExpire: accessExpire,
		Token:        token,
	}, nil
}
