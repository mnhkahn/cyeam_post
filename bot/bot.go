package bot

import (
	"cyeam_post/common"
	"cyeam_post/dao"
	"cyeam_post/parser"
	"fmt"
)

type Bot interface {
	Init(parser parser.Parser, dao dao.DaoContainer)
	Start(root string)
	Debug(is_debug bool)
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
	parser parser.Parser
	dao    dao.DaoContainer
}
