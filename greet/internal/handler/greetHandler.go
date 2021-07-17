package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"gozero_single/greet/internal/logic"
	"gozero_single/greet/internal/svc"
)

func GreetHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewGreetLogic(r.Context(), ctx)
		resp, err := l.Greet()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
