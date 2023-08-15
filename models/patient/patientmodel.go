package patient

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"meg-backup-gozero/internal/types"
)

var _ PatientModel = (*customPatientModel)(nil)

type PatientFilePath struct {
	FileName string `json:"fileName"`
}
type ChartDataInfo struct {
	Date  string `json:"date"`
	Sex   string `json:"sex"`
	Count int64  `json:"count"`
}

type (
	// PatientModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPatientModel.
	PatientModel interface {
		patientModel
		FindPatientInfoByCode(ctx context.Context, code string) ([]*PatientFilePath, error)
		GetAllPatientInfoByDid(ctx context.Context, did string) ([]*types.GetAllResponse, error)
		GetAllInformation(ctx context.Context) ([]*ChartDataInfo, error)
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

func (m *customPatientModel) GetAllInformation(ctx context.Context) ([]*ChartDataInfo, error) {
	query := fmt.Sprintf("SELECT date(upload_time) AS date, sex, COUNT(*) AS count FROM patient GROUP BY date, sex ORDER BY date, sex")
	var resp []*ChartDataInfo
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customPatientModel) FindPatientInfoByCode(ctx context.Context, code string) ([]*PatientFilePath, error) {
	query := fmt.Sprintf("select patient_info.file_path from patient inner join patient_info on patient.code = patient_info.pid where code = $1 and patient_info.file_path is not null")
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

func (m *customPatientModel) GetAllPatientInfoByDid(ctx context.Context, did string) ([]*types.GetAllResponse, error) {
	query := fmt.Sprintf("select p.id, p.did, p.name, p.sex, p.age, p.upload_time, p.code, COALESCE(pi.file_path, 'null') AS file_path from patient as p left outer join patient_info as pi on p.code = pi.pid and pi.id = (select max (id) from patient_info where pid = p.code) where did= $1 order by p.id;")
	var resp []*types.GetAllResponse
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, did)
	switch err {
	case nil:
		//fmt.Println(resp)
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
