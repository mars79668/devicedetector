package device

import (
	"io/fs"
	"path/filepath"
)

const ParserNameCamera = `camera`
const FixtureFileCamera = `cameras.yml`

func init() {
	RegDeviceParser(ParserNameCamera,
		func(fsys fs.FS, dir string) DeviceParser {
			return NewCamera(fsys, filepath.Join(dir, FixtureFileCamera))
		})
}

func NewCamera(fsys fs.FS, fileName string) *Camera {
	c := &Camera{}
	if err := c.Load(fsys, fileName); err != nil {
		return nil
	}
	return c
}

// Device parser for camera detection
type Camera struct {
	DeviceParserAbstract
}

func (c *Camera) Parse(ua string) *DeviceMatchResult {
	if !c.PreMatch(ua) {
		return nil
	}
	return c.DeviceParserAbstract.Parse(ua)
}
