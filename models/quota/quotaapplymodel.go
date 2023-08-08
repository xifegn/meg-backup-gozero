package quota

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ QuotaApplyModel = (*customQuotaApplyModel)(nil)

type (
	// QuotaApplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customQuotaApplyModel.
	QuotaApplyModel interface {
		quotaApplyModel
	}

	customQuotaApplyModel struct {
		*defaultQuotaApplyModel
	}
)

// NewQuotaApplyModel returns a model for the database table.
func NewQuotaApplyModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) QuotaApplyModel {
	return &customQuotaApplyModel{
		defaultQuotaApplyModel: newQuotaApplyModel(conn, c, opts...),
	}
}
