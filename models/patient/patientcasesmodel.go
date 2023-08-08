package patient

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PatientCasesModel = (*customPatientCasesModel)(nil)

type (
	// PatientCasesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPatientCasesModel.
	PatientCasesModel interface {
		patientCasesModel
	}

	customPatientCasesModel struct {
		*defaultPatientCasesModel
	}
)

// NewPatientCasesModel returns a model for the database table.
func NewPatientCasesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PatientCasesModel {
	return &customPatientCasesModel{
		defaultPatientCasesModel: newPatientCasesModel(conn, c, opts...),
	}
}
