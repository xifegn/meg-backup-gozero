package doctor

import (
	"context"
	"fmt"
	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDataLogic {
	return &GetDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDataLogic) GetData(req *types.GetDataRequest) (resp []*types.GetDataResponse, err error) {
	res, err := l.svcCtx.DoctorModel.GetDataByUsername(l.ctx, req.Username)
	if err != nil {
		return nil, err
	}
	//fmt.Println(res[0].CaseCount)
	data := make([]*types.GetDataResponse, 0, 4)
	data = append(data, &types.GetDataResponse{
		CardValue: fmt.Sprintf("人数：%d人", res[0].PatientCount),
		Name:      "患者人数统计",
		Value:     res[0].PatientCount,
	})
	data = append(data, &types.GetDataResponse{
		CardValue: fmt.Sprintf("次数：%d次", res[0].FileCount),
		Name:      "模型调用次数",
		Value:     res[0].FileCount,
	})
	data = append(data, &types.GetDataResponse{
		CardValue: fmt.Sprintf("个数：%d个", res[0].CaseCount),
		Name:      "诊断书数量",
		Value:     res[0].CaseCount,
	})
	return data, nil
}
