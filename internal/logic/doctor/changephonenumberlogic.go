package doctor

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
