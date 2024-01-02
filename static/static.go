// Package static содержит статические страницы встроенного сайта.
package static

import (
	"embed"
	"io/fs"
	"os"
)

//go:embed */*
var embedFS embed.FS

// New возвращает втроенную ФС (если root =="") или заданную (иначе).
func New(root string) (fs.FS, error) {
	var subtree fs.FS
	if root != "" {
		// take real filesystem
		subtree = os.DirFS(root)
	} else {
		// take embedded fs
		subtree = embedFS
	}
	return subtree, nil
}
