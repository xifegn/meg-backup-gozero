package doctor

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"
	"meg-backup-gozero/models/patient"
	"time"
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
	res, _ := l.svcCtx.DoctorModel.HaveCase(l.ctx, req.Cid)
	if res == -100 {
		fmt.Println("你好哈哈哈")
		newCase := patient.PatientCases{
			Body:       sql.NullString{req.Body, true},
			CreateTime: time.Now(),
			Cid:        req.Cid,
		}
		_, err := l.svcCtx.PatientCasesModel.Insert(l.ctx, &newCase)
		if err != nil {
			return nil, err
		}
		return &types.InsertCaseResponse{Message: "Insert success"}, nil
	}
	//fmt.Println("----------------res", res)
	updateCase := patient.PatientCases{
		Id:         res,
		Body:       sql.NullString{req.Body, true},
		CreateTime: time.Now(),
		Cid:        req.Cid,
	}
	err = l.svcCtx.PatientCasesModel.Update(l.ctx, &updateCase)
	if err != nil {
		return nil, err
	}
	return &types.InsertCaseResponse{Message: "Update success"}, nil
}
