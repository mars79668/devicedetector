package client

import (
	"io/fs"
	"path/filepath"
)

const ParserNameMediaPlayer = `mediaplayer`
const FixtureFileMediaPlayer = `mediaplayers.yml`

func init() {
	RegClientParser(ParserNameMediaPlayer,
		func(fsys fs.FS, dir string) ClientParser {
			return NewMediaPlayer(fsys, filepath.Join(dir, FixtureFileMediaPlayer))
		})
}

func NewMediaPlayer(fsys fs.FS, fileName string) *MediaPlayer {
	c := &MediaPlayer{}
	c.ParserName = ParserNameMediaPlayer
	if err := c.Load(fsys, fileName); err != nil {
		return nil
	}
	return c
}

// Client parser for mediaplayer detection
type MediaPlayer struct {
	ClientParserAbstract
}
