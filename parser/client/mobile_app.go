package client

import (
	"io/fs"
	"path/filepath"
)

const ParserNameMobileApp = `mobile app`
const FixtureFileMobileApp = `mobile_apps.yml`

func init() {
	RegClientParser(ParserNameMobileApp,
		func(fsys fs.FS, dir string) ClientParser {
			return NewMobileApp(fsys, filepath.Join(dir, FixtureFileMobileApp))
		})
}

func NewMobileApp(fsys fs.FS, fileName string) *MobileApp {
	c := &MobileApp{}
	c.ParserName = ParserNameMobileApp
	if err := c.Load(fsys, fileName); err != nil {
		return nil
	}
	return c
}

// Client parser for mobile app detection
type MobileApp struct {
	ClientParserAbstract
}
