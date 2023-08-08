package doctor

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"meg-backup-gozero/internal/logic/doctor"
	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"
	"meg-backup-gozero/response"
	"net/http"
)

func GetPatientFilePathHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPatientFilePathRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := doctor.NewGetPatientFilePathLogic(r.Context(), svcCtx)
		resp, err := l.GetPatientFilePath(&req)
		response.Response(w, resp, err)
	}
}
