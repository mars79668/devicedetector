package client

import (
	"io/fs"
	"path/filepath"
)

const ParserNamePim = `pim`
const FixtureFilePim = `pim.yml`

func init() {
	RegClientParser(ParserNamePim,
		func(fsys fs.FS, dir string) ClientParser {
			return NewPim(fsys, filepath.Join(dir, FixtureFilePim))
		})
}

func NewPim(fsys fs.FS, fileName string) *Pim {
	c := &Pim{}
	c.ParserName = ParserNamePim
	if err := c.Load(fsys, fileName); err != nil {
		return nil
	}
	return c
}

// Client parser for pim (personal information manager) detection
type Pim struct {
	ClientParserAbstract
}
