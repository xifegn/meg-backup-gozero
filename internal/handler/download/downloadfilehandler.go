package download

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"meg-backup-gozero/internal/logic/download"
	"meg-backup-gozero/internal/svc"
	"meg-backup-gozero/internal/types"
	"meg-backup-gozero/response"
	"net/http"
)

func DownloadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadFileRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := download.NewDownloadFileLogic(r.Context(), svcCtx)
		resp, err := l.DownloadFile(&req)
		response.Response(w, resp, err)
	}
}
