package preidct

import (
	"meg-backup-gozero/internal/logic/preidct"
	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/response"
	"net/http"
)

func SmUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := preidct.NewSmUploadLogic(r.Context(), svcCtx)
		resp, err := l.SmUpload()
		response.Response(w, resp, err)
	}
}
