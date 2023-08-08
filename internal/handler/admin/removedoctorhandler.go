package admin

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"meg-backup-gozero/internal/logic/admin"
	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"
	"meg-backup-gozero/response"
	"net/http"
)

func RemoveDoctorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveDoctorRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := admin.NewRemoveDoctorLogic(r.Context(), svcCtx)
		resp, err := l.RemoveDoctor(&req)
		response.Response(w, resp, err)
	}
}
