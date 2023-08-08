package preidct

import (
	"meg-backup-gozero/internal/logic/preidct"
	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/response"
	"net/http"
)

func GetPidHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := preidct.NewGetPidLogic(r.Context(), svcCtx)
		resp, err := l.GetPid()
		response.Response(w, resp, err)
	}
}
