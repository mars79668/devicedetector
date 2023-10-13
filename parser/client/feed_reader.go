package client

import (
	"io/fs"
	"path/filepath"
)

const ParserNameFeedReader = `feed reader`
const FixtureFileFeedReader = `feed_readers.yml`

func init() {
	RegClientParser(ParserNameFeedReader,
		func(fsys fs.FS, dir string) ClientParser {
			return NewFeedReader(fsys, filepath.Join(dir, FixtureFileFeedReader))
		})
}

func NewFeedReader(fsys fs.FS, fileName string) *FeedReader {
	c := &FeedReader{}
	c.ParserName = ParserNameFeedReader
	if err := c.Load(fsys, fileName); err != nil {
		return nil
	}
	return c
}

// Client parser for feed reader detection
type FeedReader struct {
	ClientParserAbstract
}

//type FeedReader struct {
//	RegularName `yaml:",inline" json:",inline"`
//	Version     string `yaml:"version" json:"version"`
//	Url         string `yaml:"url" json:"url"`
//	Type        string `yaml:"type" json:"type"`
//}
