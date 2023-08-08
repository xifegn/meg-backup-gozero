package quota

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ QuotaModel = (*customQuotaModel)(nil)

type (
	// QuotaModel is an interface to be customized, add more methods here,
	// and implement the added methods in customQuotaModel.
	QuotaModel interface {
		quotaModel
	}

	customQuotaModel struct {
		*defaultQuotaModel
	}
)

// NewQuotaModel returns a model for the database table.
func NewQuotaModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) QuotaModel {
	return &customQuotaModel{
		defaultQuotaModel: newQuotaModel(conn, c, opts...),
	}
}
