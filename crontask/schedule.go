package crontask

import (
	"log"
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
		// log.Printf("%+v", uPlugin)
		p, err := plugin.GetPlugin(uPlugin.PluginType)
		if err != nil {
			return err
		}

		shouldNotice, err := p.Execute(uPlugin)
		if err != nil {
			return err
		}

		// shouldNotice = true
		if shouldNotice {
			formid, err := store.PopEnergy(uPlugin.UserID)
			if err != nil {
				return err
			}

			msg := wechat.NewTemplateMsg(
				uPlugin.UserID,
				p.GetTemplateMsgID(),
				formid,
				uPlugin.Values,
			)
			emphasis := p.GetEmphasisID()
			if emphasis == "" {
				emphasis = "1"
			}
			msg.SetEmphasis(emphasis)
			msg.SetPage(p.GetPage())

			log.Printf("send a %s(%s) message to %s", uPlugin.PluginType, uPlugin.PluginID, uPlugin.UserID)
			return wechat.SendMsg(msg)
		}

		return nil
	}

	// crontab
	c := crontab.New()
	c.AddFunc(
		"0 * * * * *",
		func() {
			curtime := time.Now().Round(time.Minute).Unix()
			store.FetchTasks(curtime, pluginExecutor)
		},
	)
	c.Start()
}
