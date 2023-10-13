package device

import (
	"io/fs"
	"path/filepath"
)

const ParserNameConsole = `console`
const FixtureFileConsole = `consoles.yml`

func init() {
	RegDeviceParser(ParserNameConsole,
		func(fsys fs.FS, dir string) DeviceParser {
			return NewConsole(fsys, filepath.Join(dir, FixtureFileConsole))
		})
}

func NewConsole(fsys fs.FS, fileName string) *Console {
	c := &Console{}
	if err := c.Load(fsys, fileName); err != nil {
		return nil
	}
	return c
}

// Device parser for console detection
type Console struct {
	DeviceParserAbstract
}

func (c *Console) Parse(ua string) *DeviceMatchResult {
	if !c.PreMatch(ua) {
		return nil
	}
	return c.DeviceParserAbstract.Parse(ua)
}
