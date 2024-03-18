package user

import (
	"encoding/json"
	"github.com/k6mil6/backend-test-task/internal/httpgateway/http-server/model"
	"github.com/valyala/fasthttp"
)

func GetKlines(ctx *fasthttp.RequestCtx) {
	body := ctx.Request.Body()

	var req model.KlineRequest

	if err := json.Unmarshal(body, &req); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	//TODO: send to grpc service

	ctx.SetStatusCode(fasthttp.StatusOK)
}
