package logs

import (
	"github.com/astaxie/beego/logs"
)

var Log = logs.NewLogger(10000)

func init() {
	// Log.SetLogger("file", `{"filename":"test.log"}`)
	Log.SetLogger("cylog", `{"username":"lichao0407@gmail.com","password":"selinai5","host":"smtp.gmail.com:587","sendTos":["cyeamtest@gmail.com"]}`)
}
