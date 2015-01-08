package bot

import (
	"cyeam_post/common"
	"cyeam_post/dao"
	"cyeam_post/models"
	"cyeam_post/parser"
	// "fmt"
	. "cyeam_post/logs"
	"reflect"
	"strings"
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
	this.whitelist = []string{"cyeam.com"}
}

func (this *CyBot) Start(root string) {
	root = this.GetUrl(root)
	res := make(map[string]*models.Post, 0)
	Q := []string{root}
	Q_next := []string{}

	// var a = 0
	for len(Q) > 0 || len(Q_next) > 0 {
		// fmt.Println("***********", Q)
		for len(Q) != 0 {
			u := Q[0]
			Q = Q[1:]
			if _, ok := res[u]; !ok { // 过滤掉抓取过的网页
				post, next_urls := this.new(u)
				for _, next_url := range next_urls {
					exist := false
					for _, white := range this.whitelist {
						if strings.Index(next_url, white) >= 0 {
							exist = true
							break
						}
					}
					// fmt.Println(exist, next_url)
					if exist {
						Q_next = append(Q_next, next_url)
					}
				}
				res[u] = post
			}
		}
		Log.Info("%v | %s", Q_next, res)
		Q = Q_next
		Q_next = []string{}
	}
}

func (this *CyBot) new(root string) (*models.Post, []string) {
	post := new(models.Post)
	post.Link = root
	next_urls, _ := this.parser.ParseHtml(post)
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
