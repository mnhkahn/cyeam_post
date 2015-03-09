package bot

import (
	"cyeam_post/conf"
	"strings"
)

type CyBot struct {
	BotBase
}

func (this *CyBot) Prepare() {
	this.initDaoParser("duoshuo", conf.String("duoshuo.config"), "CyParser")
	this.limit = conf.DefaultInt("parse.maxcount", DEFAULT_PARSE_LIMIT)
	this.log_level = conf.Int("log.level")
	this.whitelist = strings.Split(conf.String("parse.whitelist"), ";")
	this.BotBase.Prepare()
}

func init() {
	Register("CyBot", &CyBot{})
}
