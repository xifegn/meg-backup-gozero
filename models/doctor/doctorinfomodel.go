package doctor

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DoctorInfoModel = (*customDoctorInfoModel)(nil)

type (
	// DoctorInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDoctorInfoModel.
	DoctorInfoModel interface {
		doctorInfoModel
	}

	customDoctorInfoModel struct {
		*defaultDoctorInfoModel
	}
)

// NewDoctorInfoModel returns a model for the database table.
func NewDoctorInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) DoctorInfoModel {
	return &customDoctorInfoModel{
		defaultDoctorInfoModel: newDoctorInfoModel(conn, c, opts...),
	}
}
