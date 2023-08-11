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

var _ DoctorModel = (*customDoctorModel)(nil)

type getDataContainer struct {
	FileCount    int64 `json:"file_count"`
	CaseCount    int64 `json:"case_count"`
	PatientCount int64 `json:"patient_count"`
}

type DoctorContainer struct {
	Id          int64  `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	IsAdmin     int64  `json:"is_admin"`
	Name        string `json:"name"`
	Phonenumber string `json:"phonenumber"`
	UploadTime  string `json:"upload_time"`
}

//	type DoctorInfoContainer struct {
//		Id              int64  `json:"id"`
//		Email           string `json:"email"`
//		NickName        string `json:"nickname"`
//		Regions         string `json:"regions"`
//		SelfInformation string `json:"self_information"`
//		SecretKey       string `json:"secret_key"`
//		Did             string `json:"did"`
//	}
type (
	// DoctorModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDoctorModel.
	DoctorModel interface {
		doctorModel
		AdminGetAllDoctor(ctx context.Context, code int64) ([]*types.AdminGetAllDoctorResponse, error)
		HaveCase(ctx context.Context, cid string) (int64, error)
		QueryCaseByCid(ctx context.Context, cid string) (string, error)
		GetDataByUsername(ctx context.Context, username string) ([]*getDataContainer, error)
		GetData(ctx context.Context) ([]*getDataContainer, error)
		FindIdByUsername(ctx context.Context, username string) (int64, error)
		HaveDoctorInfo(ctx context.Context, did string) (int64, error)
		//DoctorInfoByUsername(ctx context.Context, username string) ([]*types.GetDoctorInfoResponse, error)
		GetAllDoctorByUsername(ctx context.Context, username string) ([]*DoctorContainer, error)
		//GetAllDoctorInfoByUsername(ctx context.Context, username string) ([]*DoctorInfoContainer, error)
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

//func (m *customDoctorModel) GetAllDoctorInfoByUsername(ctx context.Context, username string) ([]*DoctorInfoContainer, error) {
//	query := fmt.Sprintf("select * from doctor_info where did = $1")
//	var resp []*DoctorInfoContainer
//	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, username)
//	switch err {
//	case nil:
//		return resp, nil
//	case sqlc.ErrNotFound:
//		fmt.Println("----------------------------------------------sdfasdfadsf")
//		return nil, errors.New("ErrNotFound")
//	default:
//		fmt.Println("-----------------------------------------------------------++++++++++++++++++++++++++++++")
//
//		return nil, errors.New("hahahaha")
//	}
//}

func (m *customDoctorModel) GetAllDoctorByUsername(ctx context.Context, username string) ([]*DoctorContainer, error) {
	query := fmt.Sprintf("select id, username, password, is_admin, name, phonenumber, upload_time from doctor where username=$1")
	var resp []*DoctorContainer
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

func (m *customDoctorModel) GetData(ctx context.Context) ([]*getDataContainer, error) {
	query := fmt.Sprintf("SELECT COUNT(DISTINCT pi.id) AS file_count, COUNT(DISTINCT c.id) AS case_count, COUNT(DISTINCT p.code) AS patient_count FROM doctor d LEFT JOIN patient p ON d.username = p.did LEFT JOIN patient_info pi ON p.code = pi.pid LEFT JOIN patient_cases c ON p.code = c.cid WHERE (pi.upload_date = CURRENT_DATE OR c.create_time = CURRENT_DATE) OR p.upload_time = CURRENT_DATE")
	var resp []*getDataContainer
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

func (m *customDoctorModel) GetDataByUsername(ctx context.Context, username string) ([]*getDataContainer, error) {
	query := fmt.Sprintf("SELECT COUNT(DISTINCT pi.id) AS file_count, COUNT(DISTINCT c.id) AS case_count, COUNT(DISTINCT p.code) AS patient_count FROM doctor d LEFT JOIN patient p ON d.username = p.did LEFT JOIN patient_info pi ON p.code = pi.pid LEFT JOIN patient_cases c ON p.code = c.cid WHERE d.username = $1 AND (pi.upload_date = CURRENT_DATE OR c.create_time = CURRENT_DATE OR p.upload_time = CURRENT_DATE)")
	var resp []*getDataContainer
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, errors.New("ErrNotFound")
	default:
		return nil, err
	}
}

func (m *customDoctorModel) QueryCaseByCid(ctx context.Context, cid string) (string, error) {
	query := fmt.Sprintf("SELECT body FROM patient_cases WHERE cid = $1")
	var resp string
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, cid)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return "sqlcErrNotFound", ErrNotFound
	default:
		return "", nil
	}
}

func (m *customDoctorModel) HaveCase(ctx context.Context, cid string) (int64, error) {
	query := fmt.Sprintf("select id from patient_cases where cid = $1")
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, cid)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return -100, errors.New("not found")
	default:
		return -100, errors.New("nonono")
	}
}

func (m *customDoctorModel) HaveDoctorInfo(ctx context.Context, did string) (int64, error) {
	query := fmt.Sprintf("select id from doctor_info where did = $1")
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, did)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return -100, errors.New("ErrNotFound")
	default:
		return -100, errors.New("hahaha")
	}
}

func (m *customDoctorModel) FindIdByUsername(ctx context.Context, username string) (int64, error) {
	query := fmt.Sprintf("select id from doctor where username = $1")
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return -100, errors.New("ErrNotFound")
	default:
		return -100, errors.New("hahaha")
	}
}

func (m *customDoctorModel) AdminGetAllDoctor(ctx context.Context, code int64) ([]*types.AdminGetAllDoctorResponse, error) {
	query := fmt.Sprint("select d.id, d.username, d.name, d.phonenumber, d.upload_time from doctor as d where is_admin=$1")
	var resp []*types.AdminGetAllDoctorResponse
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, code)
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

//
//func (m *customDoctorModel) DoctorInfoByUsername(ctx context.Context, username string) ([]*types.GetDoctorInfoRequest, error) {
//	query := fmt.Sprintf("SELECT d.username, d.phonenumber, d.id, di.secret_key, di.email, d.name, di.regions, di.self_information FROM doctor d LEFT JOIN doctor_info di ON d.username = di.did WHERE d.username = $1")
//	var resp []*types.GetDoctorInfoResponse
//	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, username)
//	switch err {
//	case nil:
//		return resp, nil
//	case sqlc.ErrNotFound:
//		return nil, err
//	default:
//		return nil, err
//	}
//}
//
