package bot

import (
	"cyeam_post/common"
	"cyeam_post/dao"
	"cyeam_post/models"
	"cyeam_post/parser"
	"fmt"
	"reflect"
)

type CyBot struct {
	common.CyeamBot
	parser.NormalParser
	parser parser.Parser
	dao    dao.DaoContainer
}

func (this *CyBot) Init(parser parser.Parser, dao dao.DaoContainer) {
	this.Name = reflect.TypeOf(this).String()
	this.parser = parser
	this.dao = dao
}

func (this *CyBot) Start(root string) {
	root = this.GetUrl(root)
	res := make(map[string]*models.Post, 0)
	Q := []string{root}
	Q_next := []string{}

	// var a = 0
	for len(Q) > 0 || len(Q_next) > 0 {
		fmt.Println(len(Q), len(Q_next))
		for len(Q) != 0 {
			u := Q[0]
			Q = Q[1:]
			if _, ok := res[u]; !ok {
				// fmt.Println("!!!!", u)
				post, next_urls := this.new(root)
				res[post.Link] = post
				Q_next = append(Q_next, next_urls...)
				res[u] = post
				// fmt.Println(Q, len(Q))
			}
		}
		Q = Q_next
		Q_next = []string{}
	}
}

func (this *CyBot) new(root string) (*models.Post, []string) {
	post := new(models.Post)
	post.Link = root
	next_urls, _ := this.parser.ParseHtml(post)
	// fmt.Println(next_urls)
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
