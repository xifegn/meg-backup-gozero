package doctor

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"meg-backup-gozero/internal/logic/doctor"
	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"
	"meg-backup-gozero/response"
	"net/http"
)

func InsertDoctorInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InsertDoctorInfoRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := doctor.NewInsertDoctorInfoLogic(r.Context(), svcCtx)
		resp, err := l.InsertDoctorInfo(&req)
		response.Response(w, resp, err)
	}
}
