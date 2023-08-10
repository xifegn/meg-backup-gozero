package patient

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PatientModel = (*customPatientModel)(nil)

type PatientFilePath struct {
	FileName string `json:"fileName"`
}

type (
	// PatientModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPatientModel.
	PatientModel interface {
		patientModel
		FindPatientInfoByCode(ctx context.Context, code string) ([]*PatientFilePath, error)
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

func (m *customPatientModel) FindPatientInfoByCode(ctx context.Context, code string) ([]*PatientFilePath, error) {
	query := fmt.Sprintf("select patient_info.file_path from patient inner join patient_info on patient.code = patient_info.pid where code = $1")
	var resp []*PatientFilePath
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, code)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

//func (m *customPatientModel) FindPatientInfoByCode(ctx context.Context, code string) ([]*PatientFilePath, error) {
//publicPatientCodeKey := fmt.Sprintf("%s%v", cachePublicPatientCodePrefix, code)
//var resp []*PatientFilePath
//err := m.QueryRowIndexCtx(ctx, &resp, publicPatientCodeKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
//	query := fmt.Sprintf("select patient_info.file_path from patient inner join patient_info on patient.code = patient_info.pid where code = $1")
//	if err := conn.QueryRowsCtx(ctx, &resp, query, code); err != nil {
//		return nil, err
//	}
//	return resp[0].Id, nil
//}, func(v interface{}) string {
//	return fmt.Sprintf("select %s from %s where id = %v limit 1", patientFilePathRows, m.table, v)
//})
//switch err {
//case nil:
//	return resp, nil
//case sqlc.ErrNotFound:
//	return nil, ErrNotFound
//default:
//	return nil, err
//}
//}
