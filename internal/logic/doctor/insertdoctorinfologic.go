package doctor

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
