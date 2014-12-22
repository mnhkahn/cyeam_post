package bot

import (
	"cyeam_post/models"
	"fmt"
)

type BotContainer interface {
	Start()
	ParseHtml(url string) (models.Post, []string, error)
	Debug(is_debug bool)
}

type Bot interface {
	NewJob(job string) (BotContainer, error)
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

func NewBot(name, job string) (BotContainer, error) {
	bot, ok := bots[name]
	if !ok {
		return nil, fmt.Errorf("bot: unknown bot_name %q", name)
	}
	return bot.NewJob(job)
}

type CyeamBot struct {
	Name    string
	Version string
	Job     string
	IsDebug bool
}
