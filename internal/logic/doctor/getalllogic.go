package doctor

import (
	"context"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllLogic {
	return &GetAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllLogic) GetAll(req *types.GetAllRequest) (resp []*types.GetAllResponse, err error) {
	res, err := l.svcCtx.PatientModel.GetAllPatientInfoByDid(l.ctx, req.Did)
	if err != nil {
		return nil, err
	}
	PatientInfo := make([]*types.GetAllResponse, 0)
	for _, item := range res {
		PatientInfo = append(PatientInfo, &types.GetAllResponse{
			Id:         item.Id,
			Did:        item.Did,
			Name:       item.Name,
			Sex:        item.Sex,
			Age:        item.Age,
			UploadTime: item.UploadTime[:10],
			Code:       item.Code,
			FilePath:   item.FilePath,
		})
	}

	return PatientInfo, nil
}
