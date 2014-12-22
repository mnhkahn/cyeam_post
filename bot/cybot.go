package bot

import (
	. "cyeam_post/bot/graph"
	"cyeam_post/models"
	// "github.com/astaxie/beego/httplib"
)

type CyBot struct {
}

func (this *CyBot) NewJob(job string) (BotContainer, error) {
	bot := new(CyBotContainer)
	bot.Name = "CyBot"
	bot.Version = "1.0"
	bot.Job = job
	return bot, nil
}

type CyBotContainer struct {
	CyeamBot
	g *Graph
}

func (this *CyBotContainer) Start() {
	g := Create()
	g.InsertVertex(*NewVertex(this.Job))
	for !g.IsEmpty() {
		url_v := g.Adj[0]
		_, next_urls, _ := this.ParseHtml(url_v.Data)
		g.DeleteVertex(url_v)
		for i := 0; i < len(next_urls); i++ {
			g.InsertVertex(*NewVertex(next_urls[i]))
		}
		g.Print()
		println("********")
	}
}

func (this *CyBotContainer) ParseHtml(url string) (models.Post, []string, error) {
	post := models.Post{}
	return post, nil, nil
}

func (this *CyBotContainer) Debug(is_debug bool) {
	this.IsDebug = is_debug
}

func init() {
	Register("cybot1.0", &CyBot{})
}
