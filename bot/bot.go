package bot

import (
	"cyeam_post/common"
	"cyeam_post/dao"
	. "cyeam_post/logs"
	"cyeam_post/models"
	"cyeam_post/parser"
	"fmt"
	"strings"
)

const (
	DEFAULT_PARSE_LIMIT = 1
	LOG_LEVEL_BOT       = 1
	LOG_LEVEL_DAO       = 2
	LOG_LEVEL_PARSER    = 4
)

type Bot interface {
	Prepare()
	Start(root string)
	Parse(root string) (*models.Post, []string)
	Save(*models.Post)
	Limit(maxcount int)
	ParseCount() int
	Version() string
	SetVersion(v string)
}

var bots = make(map[string]Bot)

func Register(name string, bot Bot) {
	bot.SetVersion(name)
	if bot == nil {
		panic("bot: Register bot is nil")
	}
	if _, ok := bots[name]; ok {
		panic("bot: Register called twice for bot" + name)
	}
	bots[name] = bot
}

func NewBot(name string) (Bot, error) {
	bot, ok := bots[name]
	if !ok {
		return nil, fmt.Errorf("bot: unknown bot_name %q", name)
	}
	return bot, nil
}

type BotBase struct {
	common.CyeamBot
	limit       int
	parse_count int
	log_level   int
	whitelist   []string
	parser      parser.Parser
	dao         dao.DaoContainer
}

func (this *BotBase) CountOne() {
	this.parse_count++
}

func (this *BotBase) initDaoParser(dao_name, dao_host, parser_type string) {
	Dao, err := dao.NewDao(dao_name, dao_host)
	if err != nil {
		panic(err)
	}
	Parser, err := parser.NewParser(parser_type)
	if err != nil {
		panic(err)
	}
	this.parser = Parser
	this.dao = Dao
}

func (this *BotBase) Prepare() {
	if this.log_level&LOG_LEVEL_BOT^LOG_LEVEL_BOT == 0 {
		this.Debug(true)
	}
	if this.log_level&LOG_LEVEL_DAO^LOG_LEVEL_DAO == 0 {
		this.dao.Debug(true)
	}
	if this.log_level&LOG_LEVEL_PARSER^LOG_LEVEL_PARSER == 0 {
		this.parser.Debug(true)
	}
}

func (this *BotBase) Start(root string) {
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
				post, next_urls := this.Parse(u)
				if post != nil {
					// If got nothing by parsing, skip it
					if post.Description != "" {
						if this.IsDebug {
							Log.Info("Parse %s success", post.Link)
						}
						this.Save(post)
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

func (this *BotBase) Parse(root string) (*models.Post, []string) {
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

func (this *BotBase) Save(post *models.Post) {
	this.dao.AddPost(post)
}

func (this *BotBase) ParseCount() int {
	return this.parse_count
}

func (this *BotBase) Version() string {
	return this.Name
}

func (this *BotBase) SetVersion(v string) {
	this.Name = v
}

func (this *BotBase) Limit(maxcount int) {
	this.limit = maxcount
}

func (this *BotBase) Debug(is_debug bool) {
	this.IsDebug = is_debug
}
