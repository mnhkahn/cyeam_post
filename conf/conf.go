package conf

import (
	"github.com/astaxie/beego/config"
)

var AppConfig config.ConfigContainer

func init() {
	var err error
	AppConfig, err = config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
}

func String(key string) string {
	return AppConfig.String(key)
}
func Strings(key string) []string {
	return AppConfig.Strings(key)
}
func Int(key string) int {
	i, _ := AppConfig.Int(key)
	return i
}
func Int64(key string) (int64, error) {
	return AppConfig.Int64(key)
}
func Bool(key string) bool {
	b, _ := AppConfig.Bool(key)
	return b
}
func Float(key string) (float64, error) {
	return AppConfig.Float(key)
}
func DefaultString(key string, defaultval string) string {
	return AppConfig.DefaultString(key, defaultval)
}
func DefaultStrings(key string, defaultval []string) []string {
	return AppConfig.DefaultStrings(key, defaultval)
}
func DefaultInt(key string, defaultval int) int {
	return AppConfig.DefaultInt(key, defaultval)
}
func DefaultInt64(key string, defaultval int64) int64 {
	return AppConfig.DefaultInt64(key, defaultval)
}
func DefaultBool(key string, defaultval bool) bool {
	return AppConfig.DefaultBool(key, defaultval)
}
func DefaultFloat(key string, defaultval float64) float64 {
	return AppConfig.DefaultFloat(key, defaultval)
}
