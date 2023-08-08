package patient

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PatientModel = (*customPatientModel)(nil)

type (
	// PatientModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPatientModel.
	PatientModel interface {
		patientModel
	}

	customPatientModel struct {
		*defaultPatientModel
	}
)

// NewPatientModel returns a model for the database table.
func NewPatientModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PatientModel {
	return &customPatientModel{
		defaultPatientModel: newPatientModel(conn, c, opts...),
	}
}
