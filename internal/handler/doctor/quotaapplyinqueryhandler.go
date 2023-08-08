package doctor

import (
	"meg-backup-gozero/internal/logic/doctor"
	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/response"
	"net/http"
)

func QuotaApplyInqueryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := doctor.NewQuotaApplyInqueryLogic(r.Context(), svcCtx)
		resp, err := l.QuotaApplyInquery()
		response.Response(w, resp, err)
	}
}
