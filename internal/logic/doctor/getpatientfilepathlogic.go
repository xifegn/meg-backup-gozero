package doctor

import (
	"context"
	"errors"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPatientFilePathLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPatientFilePathLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPatientFilePathLogic {
	return &GetPatientFilePathLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPatientFilePathLogic) GetPatientFilePath(req *types.GetPatientFilePathRequest) (resp []*types.GetPatientFilePathResponse, err error) {
	res, err := l.svcCtx.PatientModel.FindPatientInfoByCode(l.ctx, req.Code)
	if err != nil {
		return nil, errors.New("Not found")
	}
	PatientFilePathInfoList := make([]*types.GetPatientFilePathResponse, 0)
	for key, item := range res {
		PatientFilePathInfoList = append(PatientFilePathInfoList, &types.GetPatientFilePathResponse{
			Key:      int64(key) + 1,
			FileName: item.FileName,
		})
	}

	return PatientFilePathInfoList, nil
}
