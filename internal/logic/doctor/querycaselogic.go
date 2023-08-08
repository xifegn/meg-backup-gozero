package doctor

import (
	"context"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCaseLogic {
	return &QueryCaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryCaseLogic) QueryCase(req *types.QueryCaseRequest) (resp *types.QueryCaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
