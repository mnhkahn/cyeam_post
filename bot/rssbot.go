package bot

// import (
// 	"cyeam_post/dao"
// 	. "cyeam_post/logs"
// 	"cyeam_post/models"
// 	"cyeam_post/parser"
// 	"encoding/xml"
// 	"github.com/astaxie/beego/httplib"
// 	"reflect"
// 	"time"
// )

// type RssBot struct {
// 	BotBase
// }

// func (this *RssBot) Prepare() {

// }

// func (this *RssBot) Init(parser parser.Parser, dao dao.DaoContainer) {
// 	this.Name = reflect.TypeOf(this).String()
// 	this.parser = parser
// 	this.dao = dao
// }

// func (this *RssBot) Start(root string) {
// 	res := RssFeed{}
// 	req := httplib.Get(root)
// 	req.SetUserAgent("Cyeambot")
// 	req.SetTimeout(5*time.Second, 5*time.Second)
// 	err := req.ToXml(&res)
// 	if err != nil {
// 		panic(err)
// 	}
// 	Log.Info("Parse complete.")
// 	for _, item := range res.Channel.Items {
// 		Log.Debug(item.Link)
// 		post := new(models.Post)
// 		post.Title = item.Title
// 		if temp_date, err := time.Parse("2006-01-02T15:04:05", string([]byte(item.PubDate)[:19])); err == nil {
// 			post.CreateTime.Time = temp_date
// 		} else if temp_date, err := time.Parse("2006-01-02 15:04:05", item.PubDate); err == nil {
// 			post.CreateTime.Time = temp_date
// 		} else {
// 			post.CreateTime.Time = time.Now()
// 		}
// 		if item.Author != "" {
// 			post.Author = item.Author
// 		} else {
// 			post.Author = res.Channel.Title
// 		}
// 		post.Detail = item.Description
// 		post.Category = item.Category
// 		post.Figure = item.Figure
// 		post.Source = root
// 		post.Link = item.Link
// 		post.ParseDate.Time = time.Now()

// 		this.parser.ParseHtml(post)
// 		this.dao.AddOrUpdate(post)
// 	}
// }

// func (this *RssBot) Debug(is_debug bool) {
// 	this.IsDebug = is_debug
// }

// func (this *RssBot) Version() string {
// 	return this.Name
// }

// func init() {
// 	Register("RssBot", &RssBot{})
// }

// type RssFeed struct {
// 	XMLName xml.Name    `xml:"rss"`
// 	Channel *RssChannel `xml:"channel"`
// }

// type RssChannel struct {
// 	XMLName       xml.Name   `xml:"channel"`
// 	Title         string     `xml:"title"`
// 	Description   string     `xml:"description"`
// 	Link          string     `xml:"link"`
// 	Language      string     `xml:"language"`
// 	PubDate       string     `xml:"pubDate"`
// 	LastBuildDate string     `xml:"lastBuildDate"`
// 	Items         []*RssItem `xml:"item"`
// }

// type RssItem struct {
// 	XMLName     xml.Name `xml:"item"`
// 	Title       string   `xml:"title"` // required
// 	Figure      string   `xml:"figure"`
// 	Link        string   `xml:"link"`        // required
// 	Description string   `xml:"description"` // required
// 	Author      string   `xml:"author,omitempty"`
// 	Category    string   `xml:"category,omitempty"`
// 	Comments    string   `xml:"comments,omitempty"`
// 	Guid        string   `xml:"guid,omitempty"`    // Id used
// 	PubDate     string   `xml:"pubDate,omitempty"` // created or updated
// 	Source      string   `xml:"source,omitempty"`
// }
