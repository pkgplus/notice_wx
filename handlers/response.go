package handlers

import (
	"github.com/kataras/iris/context"
)

type Response struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Detail  string `json:"detail,omitempty"`
}

func SendResponse(ctx context.Context, code int, msg, detail string) {
	resp := &Response{
		code,
		msg,
		detail,
	}
	ctx.StatusCode(resp.Code)
	ctx.JSON(resp)
}
