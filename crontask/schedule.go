package crontask

import (
	"time"

	crontab "github.com/robfig/cron"
	"github.com/xuebing1110/noticeplat/plugin"
	"github.com/xuebing1110/noticeplat/storage"
	"github.com/xuebing1110/noticeplat/user"
	"github.com/xuebing1110/noticeplat/wechat"
)

func Schedual(store storage.Storage) {
	// plugin executor
	var pluginExecutor = func(uPlugin *user.UserPlugin) error {
		p, err := plugin.GetPlugin(uPlugin.PluginType)
		if err != nil {
			return err
		}

		shouldNotice, err := p.Execute(uPlugin)
		if err != nil {
			return err
		}

		formid, err := store.PopEnergy(uPlugin.UserID)
		if err != nil {
			return err
		}

		if shouldNotice {
			msg := wechat.NewTemplateMsg(
				uPlugin.UserID,
				p.GetTemplateMsgID(),
				formid,
				uPlugin.Values,
			)

			return wechat.SendMsg(msg)
		}

		return nil
	}

	// crontab
	c := crontab.New()
	c.AddFunc(
		"@every 1m",
		func() {
			curtime := time.Now().Round(time.Minute).Unix()
			store.FetchTasks(curtime, pluginExecutor)
		},
	)
	c.Start()
}
