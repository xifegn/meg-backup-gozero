package doctor

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DoctorModel = (*customDoctorModel)(nil)

type (
	// DoctorModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDoctorModel.
	DoctorModel interface {
		doctorModel
	}

	customDoctorModel struct {
		*defaultDoctorModel
	}
)

// NewDoctorModel returns a model for the database table.
func NewDoctorModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DoctorModel {
	return &customDoctorModel{
		defaultDoctorModel: newDoctorModel(conn, c, opts...),
	}
}
