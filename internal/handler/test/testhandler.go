package test

import (
	"meg-backup-gozero/internal/logic/test"
	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/response"
	"net/http"
)

func TestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := test.NewTestLogic(r.Context(), svcCtx)
		resp, err := l.Test()
		response.Response(w, resp, err)
	}
}
