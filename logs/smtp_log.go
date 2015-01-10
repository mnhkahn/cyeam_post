package logs

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var Log = logs.NewLogger(10000)

func init() {
	if beego.AppConfig.String("runmode") == "dev" {
		Log.SetLogger("console", `{"level":8}`)
	} else {
		Log.SetLogger("cylog", `{"username":"lichao0407@gmail.com","password":"selinai5","host":"smtp.gmail.com:587","sendTos":["cyeamtest@gmail.com"]}`)
	}
}
