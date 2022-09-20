package handler

import (
	"fmt"
	"net/http"

	"zero-strange/service/broker/api/internal/logic"
	"zero-strange/service/broker/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func getuserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path, "PATH")
		l := logic.NewGetuserLogic(r.Context(), svcCtx)
		resp, err := l.Getuser(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
