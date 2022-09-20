package handler

import (
	"net/http"

	"zero-strange/service/broker/api/internal/logic"
	"zero-strange/service/broker/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func getusersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetusersLogic(r.Context(), svcCtx)
		resp, err := l.Getusers()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
