package doctor

import (
	"context"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveFilePathLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveFilePathLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveFilePathLogic {
	return &RemoveFilePathLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveFilePathLogic) RemoveFilePath(req *types.RemoveFilePathRequest) (resp *types.RemoveFilePathResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
