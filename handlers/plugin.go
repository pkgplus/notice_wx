package handlers

import (
	"net/http"
	"sort"

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

	uid := ctx.Values().GetString(CONTEXT_OPENID_TAG)
	// unionid := ctx.Values().GetString(CONTEXT_UNION_TAG)
	if uid == "" {
		SendResponse(ctx, http.StatusInternalServerError, "read openid failed from context", "")
		return
	}

	userplugin := new(user.UserPlugin)
	err := ctx.ReadJSON(userplugin)
	if err != nil {
		SendResponse(ctx, http.StatusBadRequest, "parse to json failed", err.Error())
		return
	}
	userplugin.UserID = uid
	err = userplugin.CronSetting.Init()
	if err != nil {
		SendResponse(ctx, http.StatusBadRequest, "cron setting init failed", err.Error())
		return
	}

	err = store.AddUserPlugin(userplugin)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "add user plugin failed", err.Error())
		return
	}

	ctx.JSON(&Response{Message: "OK"})
}

func DeleteUserPlugin(ctx context.Context) {
	// storage
	store, ok := GetStorage(ctx)
	if !ok {
		SendResponse(ctx, http.StatusInternalServerError, "context exception for getting storage", "")
		return
	}

	uid := ctx.Values().GetString(CONTEXT_OPENID_TAG)
	// unionid := ctx.Values().GetString(CONTEXT_UNION_TAG)
	if uid == "" {
		SendResponse(ctx, http.StatusInternalServerError, "read openid failed from context", "")
		return
	}

	pluginid := ctx.Params().Get("pluginid")
	err := store.DelUserPlugin(uid, pluginid)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "delete user plugin failed", err.Error())
		return
	}

	ctx.JSON(&Response{Message: "OK"})
}

func ListUserPlugins(ctx context.Context) {
	// storage
	store, ok := GetStorage(ctx)
	if !ok {
		SendResponse(ctx, http.StatusInternalServerError, "context exception for getting storage", "")
		return
	}

	uid := ctx.Values().GetString(CONTEXT_OPENID_TAG)
	// unionid := ctx.Values().GetString(CONTEXT_UNION_TAG)
	if uid == "" {
		SendResponse(ctx, http.StatusInternalServerError, "read openid failed from context", "")
		return
	}

	ups, err := store.ListUserPlugins(uid)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "delete user plugin failed", err.Error())
		return
	}

	sort.Slice(ups, func(i, j int) bool {
		return ups[i].CreateTime < ups[j].CreateTime
	})

	ctx.JSON(ups)
}

func GetUserPlugin(ctx context.Context) {
	// storage
	store, ok := GetStorage(ctx)
	if !ok {
		SendResponse(ctx, http.StatusInternalServerError, "context exception for getting storage", "")
		return
	}

	uid := ctx.Values().GetString(CONTEXT_OPENID_TAG)
	// unionid := ctx.Values().GetString(CONTEXT_UNION_TAG)
	if uid == "" {
		SendResponse(ctx, http.StatusInternalServerError, "read openid failed from context", "")
		return
	}

	pluginid := ctx.Params().Get("pluginid")
	up, err := store.GetUserPlugin(uid, pluginid)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "get user plugin failed", err.Error())
		return
	}

	ctx.JSON(up)
}
