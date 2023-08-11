package quota

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"meg-backup-gozero/internal/types"
)

var _ QuotaApplyModel = (*customQuotaApplyModel)(nil)

type (
	// QuotaApplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customQuotaApplyModel.
	QuotaApplyModel interface {
		quotaApplyModel
		GetQuotaApply(ctx context.Context) ([]*types.QuotaApplyInqueryResponse, error)
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

func (m *customQuotaApplyModel) GetQuotaApply(ctx context.Context) ([]*types.QuotaApplyInqueryResponse, error) {
	query := fmt.Sprint("SELECT qa.id, qa.amount, qa.username, qa.created_at, COALESCE (q.amount, 100) AS quota_amount FROM quota_apply qa LEFT JOIN quota q ON qa.username = q.username AND q.created_at::date = CURRENT_DATE")
	var resp []*types.QuotaApplyInqueryResponse
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, errors.New("ErrNotFound")
	default:
		return nil, errors.New("")
	}
}
