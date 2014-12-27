package bot

// import (
// 	. "cyeam_post/bot/graph"
// 	"cyeam_post/common"
// 	"cyeam_post/dao"
// 	"cyeam_post/models"
// 	"cyeam_post/parser"
// 	"reflect"
// 	// "github.com/astaxie/beego/httplib"
// )

// type CyBot struct {
// 	common.CyeamBot
// 	parser parser.Parser
// 	dao    dao.Dao
// }

// func (this *CyBot) Init(parser parser.Parser, dao dao.Dao) {
// 	this.Name = reflect.TypeOf(this).String()
// 	this.parser = parser
// 	this.dao = dao
// 	// var err error
// 	// this.parser, err = parser.NewParser(this.ParserName)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// }

// func (this *CyBot) Start(root string) {
// 	g := Create()
// 	g.InsertVertex(*NewVertex(this.Root))
// 	for !g.IsEmpty() {
// 		url_v := g.Adj[0]
// 		_, next_urls, _ := this.ParseHtml(url_v.Data)
// 		g.DeleteVertex(url_v)
// 		for i := 0; i < len(next_urls); i++ {
// 			g.InsertVertex(*NewVertex(next_urls[i]))
// 		}
// 		g.Print()
// 		println("********")
// 	}
// }

// func (this *CyBot) ParseHtml(url string) (models.Post, []string, error) {
// 	post := models.Post{}
// 	return post, nil, nil
// }

// func (this *CyBot) Debug(is_debug bool) {
// 	this.IsDebug = is_debug
// }

// func (this *CyBot) Version() string {
// 	return this.Name
// }

// func init() {
// 	Register("CyBot", &CyBot{})
// }
