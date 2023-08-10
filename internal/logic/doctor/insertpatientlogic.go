package doctor

import (
	"context"
	"meg-backup-gozero/models/patient"
	"time"

	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertPatientLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertPatientLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertPatientLogic {
	return &InsertPatientLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertPatientLogic) InsertPatient(req *types.InsertPatientRequest) (resp *types.InsertPatientResponse, err error) {
	newPatient := patient.Patient{
		Did:        req.Username,
		Name:       req.Name,
		Sex:        req.Sex,
		Age:        req.Age,
		Code:       time.Now().Format("2006010215040500")[:16],
		UploadTime: time.Now(),
	}
	res, err := l.svcCtx.PatientModel.Insert(l.ctx, &newPatient)
	if err != nil {
		return nil, err
	}
	newPatient.Id, err = res.LastInsertId()

	return &types.InsertPatientResponse{
		Message: "Insert patient Ok!",
	}, nil
}
