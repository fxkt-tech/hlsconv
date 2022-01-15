package file

import (
	"path"
	"path/filepath"
	"strings"
)

func IsFolder(s string) bool {
	return strings.HasSuffix(s, "/")
}

func ParseFilesBySuffix(path, suffix string, rectusion bool) (fs []string, err error) {
	var gpath string
	if strings.HasSuffix(path, "/") {
		gpath = path + "*"
	} else {
		gpath = path + "/*"
	}
	files, err := filepath.Glob(gpath)
	if err != nil {
		return
	}
	for _, file := range files {
		if strings.HasSuffix(file, suffix) {
			fs = append(fs, file)
		}
	}
	return
}

func FileName(s string) string {
	return strings.TrimSuffix(path.Base(s), path.Ext(s))
}

func FixWinPath(s string) string {
	return strings.ReplaceAll(s, "\\", "/")
}
