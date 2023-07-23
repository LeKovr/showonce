// Package static содержит статические страницы встроенного сайта.
package static

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
)

//go:embed *
var embedFS embed.FS

const embedRoot = "" //static/"

// New возвращает втроенную ФС (если root =="") или заданную (иначе).
func New(root string) (http.FileSystem, error) {
	var serverRoot fs.FS
	var err error
	if root != "" {
		// take real filesystem
		serverRoot = os.DirFS(root)
	} else if embedRoot != "" {
		// take embedded subdir
		serverRoot, err = fs.Sub(embedFS, embedRoot)
	} else {
		// take embedded fs
		serverRoot = embedFS
	}
	if err != nil {
		return nil, err
	}
	hfs := http.FS(serverRoot)
	return hfs, nil
}
