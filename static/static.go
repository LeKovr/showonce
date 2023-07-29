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

// New возвращает втроенную ФС (если root =="") или заданную (иначе).
func New(root string) (http.FileSystem, error) {
	var serverRoot fs.FS
	if root != "" {
		// take real filesystem
		serverRoot = os.DirFS(root)
	} else {
		// take embedded fs
		serverRoot = embedFS
	}
	hfs := http.FS(serverRoot)
	return hfs, nil
}
