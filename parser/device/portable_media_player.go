package device

import (
	"io/fs"
	"path/filepath"
)

const ParserNamePortableMediaPlayer = `portablemediaplayer`
const FixtureFilePortableMediaPlayer = `portable_media_player.yml`

func init() {
	RegDeviceParser(ParserNamePortableMediaPlayer,
		func(fsys fs.FS, dir string) DeviceParser {
			return NewPortableMediaPlayer(fsys, filepath.Join(dir, FixtureFilePortableMediaPlayer))
		})
}

func NewPortableMediaPlayer(fsys fs.FS, fileName string) *PortableMediaPlayer {
	p := &PortableMediaPlayer{}
	if err := p.Load(fsys, fileName); err != nil {
		return nil
	}
	return p
}

// Device parser for portable media player detection
type PortableMediaPlayer struct {
	DeviceParserAbstract
}

func (p *PortableMediaPlayer) Parse(ua string) *DeviceMatchResult {
	if !p.PreMatch(ua) {
		return nil
	}
	return p.DeviceParserAbstract.Parse(ua)
}
