package admin

import (
	"meg-backup-gozero/internal/logic/admin"
	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/response"
	"net/http"
)

func GetAllDataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := admin.NewGetAllDataLogic(r.Context(), svcCtx)
		resp, err := l.GetAllData()
		response.Response(w, resp, err)
	}
}
