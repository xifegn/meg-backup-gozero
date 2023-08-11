package doctor

import (
	"context"
	"fmt"
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
	res, err := l.svcCtx.PatientInfoModel.GetPatientInfoByFilePath(l.ctx, req.FilePath)
	if err != nil {
		return nil, err
	}
	fmt.Println("--------------------------------", res[0].Id)
	err = l.svcCtx.PatientInfoModel.Delete(l.ctx, res[0].Id)
	if err != nil {
		return nil, err
	}
	return &types.RemoveFilePathResponse{
		Message: "Ok",
	}, nil
}
