package admin

import (
	"meg-backup-gozero/internal/logic/admin"
	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/response"
	"net/http"
)

func AdminGetAllDoctorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := admin.NewAdminGetAllDoctorLogic(r.Context(), svcCtx)
		resp, err := l.AdminGetAllDoctor()
		response.Response(w, resp, err)
	}
}
