package bot

import (
	"cyeam_post/conf"
	"cyeam_post/dao"
	. "cyeam_post/logs"
	"cyeam_post/models"
	"cyeam_post/parser"
	"reflect"
	"strings"
)

type CyBot struct {
	parser.NormalParser
	BotBase
}

func (this *CyBot) Prepare() {
	Dao, err := dao.NewDao("solr", conf.String("solr.host"))
	if err != nil {
		panic(err)
	}
	Parser, err := parser.NewParser("CyParser")
	if err != nil {
		panic(err)
	}

	this.Init(Parser, Dao)
	this.BotBase.Prepare()
}

func (this *CyBot) Init(parser parser.Parser, dao dao.DaoContainer) {
	this.Name = reflect.Indirect(reflect.ValueOf(this)).Type().String()
	this.limit = conf.DefaultInt("parse.maxcount", DEFAULT_PARSE_LIMIT)
	this.log_level = conf.Int("log.level")
	this.parser = parser
	this.dao = dao
	this.whitelist = strings.Split(conf.String("parse.whitelist"), ";")
}

func (this *CyBot) Start(root string) {
	res := make(map[string]*models.Post, 0)
	Q := []string{root}
	Q_next := []string{}

	for len(Q) > 0 || len(Q_next) > 0 {
		if this.parse_count > this.limit {
			Log.Notice("Exceeded the limit %d", this.limit)
			return
		}
		// Log.Debug("%v", Q)
		for len(Q) != 0 {
			if this.parse_count > this.limit {
				Log.Notice("Exceeded the limit %d", this.limit)
				return
			}
			u := Q[0]
			Q = Q[1:]
			if _, ok := res[u]; !ok { // 过滤掉抓取过的网页
				if this.IsDebug {
					Log.Info("Start parse: %s", u)
				}
				post, next_urls := this.new(u)
				if post != nil {
					// If got nothing by parsing, skip it
					if post.Description != "" {
						if this.IsDebug {
							Log.Info("Parse %s success", post.Link)
						}
						this.dao.AddPost(post)
					}
				}
				for _, next_url := range next_urls {
					exist := false
					for _, white := range this.whitelist {
						if strings.Index(next_url, white) >= 0 {
							exist = true
							break
						}
					}
					// Log.Trace("%v %s", exist, next_url)
					if exist {
						Q_next = append(Q_next, next_url)
					}
				}
				// If parse fail, it's need to save it anyway
				res[u] = post
			}
		}
		Q = Q_next
		Q_next = []string{}
	}
}

func (this *CyBot) new(root string) (*models.Post, []string) {
	this.CountOne()
	post := new(models.Post)
	post.Link = root
	next_urls, err := this.parser.ParseHtml(post)
	if err != nil {
		Log.Error(err.Error())
		return nil, nil
	}
	return post, next_urls
}

func (this *CyBot) ParseCount() int {
	return this.parse_count
}

func (this *CyBot) Version() string {
	return this.Name
}

func init() {
	Register("CyBot", &CyBot{})
}
