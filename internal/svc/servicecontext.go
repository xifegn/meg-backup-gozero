package svc

import (
	"github.com/zeromicro/go-zero/core/stores/postgres"
	"meg-backup-gozero/internal/config"
	"meg-backup-gozero/models/doctor"
	"meg-backup-gozero/models/patient"
	"meg-backup-gozero/models/quota"
)

type ServiceContext struct {
	Config            config.Config
	DoctorModel       doctor.DoctorModel
	DoctorInfoModel   doctor.DoctorInfoModel
	PatientModel      patient.PatientModel
	PatientCasesModel patient.PatientCasesModel
	PatientInfoModel  patient.PatientInfoModel
	QuotaModel        quota.QuotaModel
	QuotaApplyModel   quota.QuotaApplyModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := postgres.New(c.Postgres.DataSource)
	return &ServiceContext{
		Config:            c,
		DoctorModel:       doctor.NewDoctorModel(conn, c.CacheRedis),
		DoctorInfoModel:   doctor.NewDoctorInfoModel(conn, c.CacheRedis),
		PatientModel:      patient.NewPatientModel(conn, c.CacheRedis),
		PatientInfoModel:  patient.NewPatientInfoModel(conn, c.CacheRedis),
		PatientCasesModel: patient.NewPatientCasesModel(conn, c.CacheRedis),
		QuotaModel:        quota.NewQuotaModel(conn, c.CacheRedis),
		QuotaApplyModel:   quota.NewQuotaApplyModel(conn, c.CacheRedis),
	}
}
