package client

import (
	"fmt"
	"io/fs"
)

var clientFactory = make(map[string]func(fs.FS, string) ClientParser, 10)

func RegClientParser(name string, f func(fs.FS, string) ClientParser) {
	clientFactory[name] = f
}

func GetClientCreater(name string) func(fs.FS, string) ClientParser {
	f, _ := clientFactory[name]
	return f
}

func NewClientParser(fsys fs.FS, dir, name string) ClientParser {
	if f, ok := clientFactory[name]; ok {
		return f(fsys, dir)
	}
	return nil
}

func NewClientParsers(fsys fs.FS, dir string, names []string) []ClientParser {
	r := make([]ClientParser, len(names))
	for i, name := range names {
		if f, ok := clientFactory[name]; ok {
			r[i] = f(fsys, dir)
		}
		if r[i] == nil {
			fmt.Printf("Client is null:" + name)
		}

	}
	return r
}
