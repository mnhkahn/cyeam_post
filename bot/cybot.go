package bot

import (
	"cyeam_post/common"
	"cyeam_post/conf"
	"cyeam_post/dao"
	. "cyeam_post/logs"
	"cyeam_post/models"
	"cyeam_post/parser"
	"reflect"
	"strings"
	// "time"
)

type CyBot struct {
	common.CyeamBot
	parser.NormalParser
	parser    parser.Parser
	dao       dao.DaoContainer
	whitelist []string
}

func (this *CyBot) Init(parser parser.Parser, dao dao.DaoContainer) {
	this.Name = reflect.TypeOf(this).String()
	this.parser = parser
	this.dao = dao
	this.whitelist = strings.Split(conf.String("parse.whitelist"), ";")
}

func (this *CyBot) Start(root string) {
	root = this.GetUrl(root)
	res := make(map[string]*models.Post, 0)
	Q := []string{root}
	Q_next := []string{}

	var i = 0
	for len(Q) > 0 || len(Q_next) > 0 {
		if i > conf.DefaultInt("parse.maxcount", 1) {
			return
		}
		// Log.Debug("%v", Q)
		for len(Q) != 0 {
			if i > conf.DefaultInt("parse.maxcount", 1) {
				return
			}
			u := Q[0]
			Q = Q[1:]
			if _, ok := res[u]; !ok { // 过滤掉抓取过的网页
				Log.Info("Start parse: %s", u)
				post, next_urls := this.new(u)
				if post != nil {
					i++
					// If got nothing by parsing, skip it
					if post.Description != "" {
						Log.Info("Parse %s success", post.Link)
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
	post := new(models.Post)
	post.Link = root
	next_urls, err := this.parser.ParseHtml(post)
	if err != nil {
		Log.Error(err.Error())
		return nil, nil
	}
	return post, next_urls
}

func (this *CyBot) Debug(is_debug bool) {
	this.IsDebug = is_debug
}

func (this *CyBot) Version() string {
	return this.Name
}

func init() {
	Register("CyBot", &CyBot{})
}
