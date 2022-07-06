package static

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
)

//go:embed *
var embedFS embed.FS

const embedRoot = "" //static"

func New(root string) (hfs http.FileSystem, err error) {
	var serverRoot fs.FS
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
		return
	}
	hfs = http.FS(serverRoot)
	return
}
