package crontask

import (
	"errors"
	"time"

	crontab "github.com/robfig/cron"
	"github.com/xuebing1110/noticeplat/storage"
	"github.com/xuebing1110/noticeplat/user"

	"github.com/xuebing1110/noticeplat/plugin/hrsign"
)

func Schedual(store storage.Storage) {
	c := crontab.New()
	c.AddFunc(
		"@every 1m",
		func() {
			curtime := time.Now().Round(time.Minute).Unix()
			store.FetchTasks(curtime, pluginExecuter)
		},
	)
	c.Start()
}

func pluginExecuter(ups *user.UserPluginSetting) error {
	switch ups.PluginType {
	case "HrSign":
		sign := &hrsign.HrSignPlugin{HrUserID: "01462834"}
		return sign.Execute(ups)
	default:
		return errors.New("unknown plugin type:" + ups.PluginType)
	}
}
