package device

import "io/fs"

var deviceFactory = make(map[string]func(fs.FS, string) DeviceParser, 10)

func RegDeviceParser(name string, f func(fs.FS, string) DeviceParser) {
	deviceFactory[name] = f
}

func GetDeviceCreater(name string) func(fs.FS, string) DeviceParser {
	f, _ := deviceFactory[name]
	return f
}

func NewDeviceParser(fsys fs.FS, dir, name string) DeviceParser {
	if f, ok := deviceFactory[name]; ok {
		return f(fsys, dir)
	}
	return nil
}

func NewDeviceParsers(fsys fs.FS, dir string, names []string) []DeviceParser {
	r := make([]DeviceParser, len(names))
	for i, name := range names {
		if f, ok := deviceFactory[name]; ok {
			r[i] = f(fsys, dir)
		}
	}
	return r
}
