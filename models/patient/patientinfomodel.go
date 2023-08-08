package patient

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PatientInfoModel = (*customPatientInfoModel)(nil)

type (
	// PatientInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPatientInfoModel.
	PatientInfoModel interface {
		patientInfoModel
	}

	customPatientInfoModel struct {
		*defaultPatientInfoModel
	}
)

// NewPatientInfoModel returns a model for the database table.
func NewPatientInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PatientInfoModel {
	return &customPatientInfoModel{
		defaultPatientInfoModel: newPatientInfoModel(conn, c, opts...),
	}
}
