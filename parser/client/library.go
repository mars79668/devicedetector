package client

import (
	"io/fs"
	"path/filepath"
)

const ParserNameLibrary = `library`
const FixtureFileLibrary = `libraries.yml`

func init() {
	RegClientParser(ParserNameLibrary,
		func(fsys fs.FS, dir string) ClientParser {
			return NewLibrary(fsys, filepath.Join(dir, FixtureFileLibrary))
		})
}

func NewLibrary(fsys fs.FS, fileName string) *Library {
	c := &Library{}
	c.ParserName = ParserNameLibrary
	if err := c.Load(fsys, fileName); err != nil {
		return nil
	}
	return c
}

// Client parser for tool & software detection
type Library struct {
	ClientParserAbstract
}
