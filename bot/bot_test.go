package bot

import (
	"testing"
)

func TestMain(t *testing.T) {
	bot, err := NewBot("cybot1.0", "http://localhost:8080")
	if err != nil {
		panic(err)
	}
	bot.Debug(true)
	bot.Start()
}
