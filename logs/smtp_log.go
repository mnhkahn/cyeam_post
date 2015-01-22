package logs

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var Log = logs.NewLogger(10000)

func init() {
	Log.EnableFuncCallDepth(true)
	if beego.AppConfig.String("runmode") == "dev" {
		if beego.AppConfig.String("log.dev.type") == "file" {
			Log.SetLogger("file", `{"filename":"bot.log"}`)
		} else {
			Log.SetLogger("console", `{"level":8}`)
		}
	} else {
		Log.SetLogger("cylog", `{"username":`+beego.AppConfig.String("email.sender")+`,"password":`+beego.AppConfig.String("email.sender.pwd")+`,"host":"smtp.gmail.com:587","sendTos":[`+beego.AppConfig.String("email.receiver")+`]}`)
	}
}
