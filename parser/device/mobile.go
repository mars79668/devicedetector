package device

import (
	"io/fs"
	"path/filepath"
)

const ParserNameMobile = `mobile`
const FixtureFileMobile = `mobiles.yml`

func init() {
	RegDeviceParser(ParserNameMobile,
		func(fsys fs.FS, dir string) DeviceParser {
			return NewMobile(fsys, filepath.Join(dir, FixtureFileMobile))
		})
}

func NewMobile(fsys fs.FS, fileName string) *Mobile {
	m := &Mobile{}
	if err := m.Load(fsys, fileName); err != nil {
		return nil
	}
	return m
}

// Device parser for mobile detection
type Mobile struct {
	DeviceParserAbstract
}
