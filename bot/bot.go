package bot

import (
	"cyeam_post/common"
	"cyeam_post/dao"
	"cyeam_post/parser"
	"fmt"
)

const (
	DEFAULT_PARSE_LIMIT = 1
	LOG_LEVEL_BOT       = 1
	LOG_LEVEL_DAO       = 2
	LOG_LEVEL_PARSER    = 4
)

type Bot interface {
	Prepare()
	Init(parser parser.Parser, dao dao.DaoContainer)
	Start(root string)
	Limit(maxcount int)
	ParseCount() int
	// Debug(is_debug bool)
	Version() string
}

var bots = make(map[string]Bot)

func Register(name string, bot Bot) {
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
	parser      parser.Parser
	dao         dao.DaoContainer
}

func (this *BotBase) CountOne() {
	this.parse_count++
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

func (this *BotBase) Debug(is_debug bool) {
	this.IsDebug = is_debug
}
