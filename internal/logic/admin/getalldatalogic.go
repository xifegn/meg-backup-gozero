package admin

import (
	"context"
	"fmt"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllDataLogic {
	return &GetAllDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllDataLogic) GetAllData() (resp []*types.GetAllDataResponse, err error) {
	res, err := l.svcCtx.DoctorModel.GetData(l.ctx)
	if err != nil {
		return nil, err
	}
	//fmt.Println(res[0].CaseCount)
	data := make([]*types.GetAllDataResponse, 0, 4)
	data = append(data, &types.GetAllDataResponse{
		CardValue: fmt.Sprintf("人数：%d人", res[0].PatientCount),
		Name:      "患者人数统计",
		Value:     res[0].PatientCount,
	})
	data = append(data, &types.GetAllDataResponse{
		CardValue: fmt.Sprintf("次数：%d次", res[0].FileCount),
		Name:      "模型调用次数",
		Value:     res[0].FileCount,
	})
	data = append(data, &types.GetAllDataResponse{
		CardValue: fmt.Sprintf("个数：%d个", res[0].CaseCount),
		Name:      "诊断书数量",
		Value:     res[0].CaseCount,
	})
	return data, nil
}
