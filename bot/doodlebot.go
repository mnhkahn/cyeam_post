package bot

type DoodleBot struct {
	RssBot
}

// func (this *DoodleBot) Start(root string) {

// }

func init() {
	Register("DoodleBot", &DoodleBot{})
}
