package quota

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ QuotaModel = (*customQuotaModel)(nil)

type QuotaContainer struct {
	Day            string `json:"day"`
	FileCount      int64  `json:"file_count"`
	DailyCallLimit int64  `json:"daily_call_limit"`
}

type AddQuota struct {
	Id     int64 `json:"id"`
	Amount int64 `json:"amount"`
}

type (
	// QuotaModel is an interface to be customized, add more methods here,
	// and implement the added methods in customQuotaModel.
	QuotaModel interface {
		quotaModel
		GetQuotaForAWeekByUsername(ctx context.Context, username string) ([]*QuotaContainer, error)
		AddQuotaByUsername(ctx context.Context, username string) ([]*AddQuota, error)
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

func (m *customQuotaModel) GetQuotaForAWeekByUsername(ctx context.Context, username string) ([]*QuotaContainer, error) {
	query := fmt.Sprintf("WITH dow AS (SELECT date_part('isodow', current_date) AS dow),dates AS (SELECT date_trunc('day', generate_series(current_date - concat(dow - 1, ' days')::interval, current_date + concat(7 - dow, ' days')::interval, '1 day'::interval)) AS day  FROM dow) SELECT d.day, coalesce(p.file_count, 0) AS file_count, coalesce(q.amount, 100) AS daily_call_limit FROM dates d LEFT JOIN (SELECT date_trunc('day', upload_date) AS day, count(*) AS file_count FROM patient_info WHERE pid IN (SELECT code FROM patient WHERE did = $1) AND upload_date BETWEEN (SELECT current_date - concat(dow - 1, ' days')::interval FROM dow) AND (SELECT current_date + concat(7 - dow, ' days')::interval FROM dow) GROUP BY day) p ON d.day = p.day LEFT JOIN (SELECT created_at::date AS day, amount FROM quota WHERE username = $1) q ON d.day = q.day ORDER BY d.day")
	var resp []*QuotaContainer
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customQuotaModel) AddQuotaByUsername(ctx context.Context, username string) ([]*AddQuota, error) {
	query := fmt.Sprintf("select id, amount from quota where created_at::date = CURRENT_DATE and username = $1")
	var resp []*AddQuota
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, errors.New("ErrNotFound")
	default:
		return nil, errors.New("hahahha")
	}
}
