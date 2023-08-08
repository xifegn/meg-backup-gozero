package doctor

import (
	"context"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertPatientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertPatientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertPatientLogic {
	return &InsertPatientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertPatientLogic) InsertPatient(req *types.InsertPatientRequest) (resp *types.InsertPatientResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
