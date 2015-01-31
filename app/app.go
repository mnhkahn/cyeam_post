package app

import (
	"cyeam_post/bot"
	. "cyeam_post/logs"
	"cyeam_post/utils"
	"time"
)

var DEFAULT_BOT_CONTROLLER *BotController

type BotController struct {
	tasks map[string]bot.Bot
}

func NewBotController() *BotController {
	return &BotController{
		tasks: make(map[string]bot.Bot),
	}
}

func (this *BotController) Add(root, bot_name string) {
	b, err := bot.NewBot(bot_name)
	if err != nil {
		panic(err)
	}
	b.Prepare()

	this.tasks[root] = b
}

func (this *BotController) Run() {
	utils.ColorLog("[INFO] Start parsing at %s...\n", time.Now().Format(time.RFC3339))
	for root, b := range this.tasks {
		utils.ColorLog("[TRAC] Start to parse [%s] with [%s]...\n", root, b.Version())
		b.Start(root)
		utils.ColorLog("[TRAC] End parsing [%s] with [%d] times\n", root, b.ParseCount())
	}
	Log.Close()
	utils.ColorLog("[SUCC] Parse successful.\n")
}

func init() {
	utils.ColorLog("[INFO] Initializing bot...\n")
	DEFAULT_BOT_CONTROLLER = NewBotController()
}

func Add(root, bot_name string) {
	utils.ColorLog("[TRAC] Uses %s to parse [%s]...\n", bot_name, root)
	DEFAULT_BOT_CONTROLLER.Add(root, bot_name)
}

func Run() {
	DEFAULT_BOT_CONTROLLER.Run()
}
