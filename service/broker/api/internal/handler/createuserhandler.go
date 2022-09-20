package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-strange/service/broker/api/internal/logic"
	"zero-strange/service/broker/api/internal/svc"
	"zero-strange/service/broker/api/internal/types"
)

func createuserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateuserLogic(r.Context(), svcCtx)
		resp, err := l.Createuser(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
