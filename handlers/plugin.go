package handlers

import (
	"net/http"

	"github.com/kataras/iris/context"
	"github.com/xuebing1110/noticeplat/user"
)

func AddUserPlugin(ctx context.Context) {
	// storage
	store, ok := GetStorage(ctx)
	if !ok {
		SendResponse(ctx, http.StatusInternalServerError, "context exception for getting storage", "")
		return
	}

	// openid := ctx.Values().GetString(CONTEXT_OPENID_TAG)
	unionid := ctx.Values().GetString(CONTEXT_UNION_TAG)
	if unionid == "" {
		SendResponse(ctx, http.StatusInternalServerError, "read unionid failed from context", "")
		return
	}

	userplugin := new(user.UserPlugin)
	err := ctx.ReadJSON(userplugin)
	if err != nil {
		SendResponse(ctx, http.StatusBadRequest, "parse to json failed", err.Error())
		return
	}
	userplugin.UnionID = unionid

	err = store.AddUserPlugin(userplugin)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "add user plugin failed", err.Error())
		return
	}

	ctx.JSON(&Response{Message: "OK"})
}
