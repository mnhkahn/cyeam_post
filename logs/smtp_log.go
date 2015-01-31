package logs

import (
	"cyeam_post/conf"
	"github.com/astaxie/beego/logs"
)

var Log = logs.NewLogger(10000)

func init() {
	Log.EnableFuncCallDepth(true)
	if conf.String("log.dev.type") == "file" {
		Log.SetLogger("file", `{"filename":"bot.log"}`)
	} else if conf.String("log.dev.type") == "console" {
		Log.SetLogger("console", `{"level":8}`)
	} else if conf.String("log.dev.type") == "log" {
		Log.SetLogger("cylog", `{"username":`+conf.String("email.sender")+`,"password":`+conf.String("email.sender.pwd")+`,"host":"smtp.gmail.com:587","sendTos":[`+conf.String("email.receiver")+`]}`)
	}
}
