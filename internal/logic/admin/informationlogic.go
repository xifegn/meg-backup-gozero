package admin

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"
	"sort"
	"time"
)

type InformationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInformationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InformationLogic {
	return &InformationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InformationLogic) Information() (resp *types.ChartDataResponse, err error) {
	res, err := l.svcCtx.PatientModel.GetAllInformation(l.ctx)
	if err != nil {
		return nil, err
	}
	//fmt.Println(res)
	maxmum := 0
	dateSexCount := make(map[string]map[string]int64)
	for _, row := range res {
		if int(row.Count) > maxmum {
			maxmum = int(row.Count)
		}
		date := row.Date
		date = date[:10]
		if _, ok := dateSexCount[date]; !ok {
			dateSexCount[date] = make(map[string]int64)
			dateSexCount[date]["male"] = 0
			dateSexCount[date]["female"] = 0
		}
		dateSexCount[date][row.Sex] = row.Count
	}
	resp = &types.ChartDataResponse{
		Max: int64(maxmum),
		Series: []types.Series{
			types.Series{
				Data: []int64{},
				Name: "男患者",
				Type: "bar",
			},
			types.Series{
				Data: []int64{},
				Name: "女患者",
				Type: "bar",
			},
		},
		XAxis: types.XAxis{
			Data: []string{},
		},
	}
	dates := make([]time.Time, 0, len(dateSexCount))
	for date := range dateSexCount {
		t, err := time.Parse("2006-01-02", date)
		if err != nil {
			return nil, err
		}
		dates = append(dates, t)
	}
	sort.SliceStable(dates, func(i, j int) bool {
		return dates[i].Before(dates[j])
	})
	for _, t := range dates {
		date := t.Format("2006-01-02")
		resp.XAxis.Data = append(resp.XAxis.Data, date)
		resp.Series[0].Data = append(resp.Series[0].Data, dateSexCount[date]["male"])
		resp.Series[1].Data = append(resp.Series[1].Data, dateSexCount[date]["female"])
	}
	return resp, nil
}
