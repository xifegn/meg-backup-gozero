package preidct

import (
	"context"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SmUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSmUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SmUploadLogic {
	return &SmUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SmUploadLogic) SmUpload() (resp *types.SmUploadResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
