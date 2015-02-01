package bot

import (
	"cyeam_post/conf"
)

type RssBot struct {
	BotBase
}

func (this *RssBot) Prepare() {
	this.initDaoParser("solr", conf.String("solr.host"), "RssParser")
	this.limit = 1
	this.log_level = conf.Int("log.level")
	this.BotBase.Prepare()
}

func init() {
	Register("RssBot", &RssBot{})
}
