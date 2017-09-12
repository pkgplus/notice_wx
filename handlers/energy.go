package handlers

import (
	"github.com/kataras/iris/context"
	"net/http"
)

func AddEnery(ctx context.Context) {
	// storage
	store, ok := GetStorage(ctx)
	if !ok {
		SendResponse(ctx, http.StatusInternalServerError, "context exception for getting storage", "")
		return
	}

	eneryMap := make(map[string]string)
	err := ctx.ReadJSON(eneryMap)
	if err != nil {
		SendResponse(ctx, http.StatusBadRequest, "parse json failed", err.Error())
		return
	}

	energy, ok := eneryMap["energy"]
	if !ok {
		SendResponse(ctx, http.StatusBadRequest, "energy is required", "")
		return
	}

	uid := ctx.Values().GetString(CONTEXT_OPENID_TAG)
	// unionid := ctx.Values().GetString(CONTEXT_UNION_TAG)
	if uid == "" {
		SendResponse(ctx, http.StatusInternalServerError, "read openid failed from context", "")
		return
	}

	err = store.AddEnergy(uid, energy)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "add energy failed", err.Error())
		return
	}

	ctx.JSON(&Response{Message: "OK"})
}

func EneryCount(ctx context.Context) {
	// storage
	store, ok := GetStorage(ctx)
	if !ok {
		SendResponse(ctx, http.StatusInternalServerError, "context exception for getting storage", "")
		return
	}

	uid := ctx.Values().GetString(CONTEXT_OPENID_TAG)
	// uid := ctx.Values().GetString(CONTEXT_UNION_TAG)
	if uid == "" {
		SendResponse(ctx, http.StatusInternalServerError, "read openid failed from context", "")
		return
	}

	// count := store.GetEnergyCount(uid)
	// if err != nil {
	// 	SendResponse(ctx, http.StatusInternalServerError, "get energy count failed", err.Error())
	// 	return
	// }

	ctx.JSON(map[string]int64{"count": store.GetEnergyCount(uid)})
}

func PopEnergy(ctx context.Context) {
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

	energy, err := store.PopEnergy(uid)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "pop one energy failed", err.Error())
		return
	}

	if energy == "" {
		ctx.JSON(&Response{http.StatusOK, "no energy to pop", ""})
	} else {
		ctx.JSON(map[string]string{"enery": energy})
	}

}
