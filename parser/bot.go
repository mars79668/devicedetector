package parser

import (
	"io/fs"
	"path/filepath"
)

var botFactory = make(map[string]func(fs.FS, string) BotParser)

func RegBotParser(name string, f func(fs.FS, string) BotParser) {
	botFactory[name] = f
}

func GetBotCreater(name string) func(fs.FS, string) BotParser {
	f, _ := botFactory[name]
	return f
}

func NewBotParser(fsys fs.FS, dir, name string) BotParser {
	if f, ok := botFactory[name]; ok {
		return f(fsys, dir)
	}
	return nil
}

const ParserNameBot = `bot`
const FixtureFileBot = `bots.yml`

func init() {
	RegBotParser(ParserNameBot,
		func(fsys fs.FS, dir string) BotParser {
			return NewBot(fsys, filepath.Join(dir, FixtureFileBot))
		})
}

func NewBot(fsys fs.FS, fileName string) *Bot {
	c := &Bot{}
	c.ParserName = ParserNameBot
	if err := c.Load(fsys, fileName); err != nil {
		return nil
	}
	return c
}

// Parses a user agent for bot information
type Bot struct {
	BotParserAbstract
}
