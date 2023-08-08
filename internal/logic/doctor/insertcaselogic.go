package doctor

import (
	"context"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertCaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertCaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertCaseLogic {
	return &InsertCaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertCaseLogic) InsertCase(req *types.InsertCaseRequest) (resp *types.InsertCaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
