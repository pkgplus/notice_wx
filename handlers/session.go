package handlers

import (
	"net/http"

	"github.com/kataras/iris/context"
	"github.com/xuebing1110/noticeplat/storage"
	"github.com/xuebing1110/noticeplat/storage/redis"
)

const (
	CONTEXT_STORAGE_TAG = "Storage"
	CONTEXT_OPENID_TAG  = "OpenID"
	CONTEXT_UNION_TAG   = "UnionID"
)

func SessionStorage(ctx context.Context) {
	var redis_store storage.Storage = redis.Client
	ctx.Values().Set(CONTEXT_STORAGE_TAG, redis_store)

	ctx.Next()
}

func GetStorage(ctx context.Context) (store storage.Storage, ok bool) {
	store, ok = ctx.Values().Get(CONTEXT_STORAGE_TAG).(storage.Storage)
	return store, ok
}

func SessionCheck(ctx context.Context) {
	// storage
	store, ok := GetStorage(ctx)
	if !ok {
		SendResponse(ctx, http.StatusInternalServerError, "context exception for getting storage", "")
		return
	}

	sid := ctx.Params().Get("sid")
	resp, err := store.QuerySession(sid)
	if err != nil {
		SendResponse(ctx, http.StatusUnauthorized, "session maybe expired", err.Error())
		return
	}

	ctx.Values().Set(CONTEXT_OPENID_TAG, resp.OpenID)
	ctx.Values().Set(CONTEXT_UNION_TAG, resp.Unionid)

	ctx.Next()
}
