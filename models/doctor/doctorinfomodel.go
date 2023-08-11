package doctor

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"meg-backup-gozero/internal/types"
)

var _ DoctorInfoModel = (*customDoctorInfoModel)(nil)

type DoctorInfoContainer struct {
	Id              int64  `json:"id"`
	Email           string `json:"email"`
	NickName        string `json:"nickname"`
	Regions         string `json:"regions"`
	SelfInformation string `json:"self_information"`
	SecretKey       string `json:"secret_key"`
	Did             string `json:"did"`
}

type (
	// DoctorInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDoctorInfoModel.
	DoctorInfoModel interface {
		doctorInfoModel
		DoctorInfoByUsername(ctx context.Context, username string) ([]*types.GetDoctorInfoResponse, error)
		GetAllDoctorInfoByUsername(ctx context.Context, username string) ([]*DoctorInfoContainer, error)
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

func (m *customDoctorInfoModel) DoctorInfoByUsername(ctx context.Context, username string) ([]*types.GetDoctorInfoResponse, error) {
	query := fmt.Sprintf("SELECT d.username, d.phonenumber, d.id, di.secret_key, di.email, d.name, di.regions, di.self_information FROM doctor d LEFT JOIN doctor_info di ON d.username = di.did WHERE d.username = $1")
	var resp []*types.GetDoctorInfoResponse
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, err
	default:
		return nil, err
	}
}

func (m *customDoctorInfoModel) GetAllDoctorInfoByUsername(ctx context.Context, username string) ([]*DoctorInfoContainer, error) {
	query := fmt.Sprintf("select * from doctor_info where did = $1")
	var resp []*DoctorInfoContainer
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		//fmt.Println("----------------------------------------------sdfasdfadsf")
		return nil, errors.New("ErrNotFound")
	default:
		//fmt.Println("-----------------------------------------------------------++++++++++++++++++++++++++++++")

		return nil, errors.New("hahahaha")
	}
}
