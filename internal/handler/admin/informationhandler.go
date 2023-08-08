package admin

import (
	"meg-backup-gozero/internal/logic/admin"
	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/response"
	"net/http"
)

func InformationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := admin.NewInformationLogic(r.Context(), svcCtx)
		resp, err := l.Information()
		response.Response(w, resp, err)
	}
}
