package patient

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PatientInfoModel = (*customPatientInfoModel)(nil)

type GetPatientId struct {
	Id int64 `json:"id"`
}

type (
	// PatientInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPatientInfoModel.
	PatientInfoModel interface {
		patientInfoModel
		GetPatientInfoByFilePath(ctx context.Context, filePath string) ([]*GetPatientId, error)
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

func (m *customPatientInfoModel) GetPatientInfoByFilePath(ctx context.Context, filePath string) ([]*GetPatientId, error) {
	query := fmt.Sprintf("SELECT p.id FROM patient_info as p WHERE file_path = $1")
	var resp []*GetPatientId
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, filePath)
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
