package device

import (
	"io/fs"
	"path/filepath"
)

const ParserNameCar = `car browser`
const FixtureFileCar = `car_browsers.yml`

func init() {
	RegDeviceParser(ParserNameCar,
		func(fsys fs.FS, dir string) DeviceParser {
			return NewCar(fsys, filepath.Join(dir, FixtureFileCar))
		})
}

func NewCar(fsys fs.FS, fileName string) *Car {
	c := &Car{}
	if err := c.Load(fsys, fileName); err != nil {
		return nil
	}
	return c
}

// Device parser for car browser detection
type Car struct {
	DeviceParserAbstract
}

func (c *Car) Parse(ua string) *DeviceMatchResult {
	if !c.PreMatch(ua) {
		return nil
	}
	return c.DeviceParserAbstract.Parse(ua)
}
