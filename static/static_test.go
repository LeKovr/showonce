package static_test

import (
	"testing"

	"github.com/LeKovr/showonce/static"
	ass "github.com/alecthomas/assert/v2"
)

func TestNewEmbed(t *testing.T) {
	got, err := static.New("html/")
	ass.NotZero(t, got, "FS not nil")
	ass.NoError(t, err, "New success")
	f, err := got.Open("index.html")
	errC := f.Close()
	ass.NoError(t, err, "Open success")
	ass.NoError(t, errC, "Close success")
}

func TestNewReal(t *testing.T) {
	got, err := static.New("html/js")
	ass.NotZero(t, got, "FS not nil")
	ass.NoError(t, err, "New success")
	f, err := got.Open("service.swagger.json")
	errC := f.Close()
	ass.NoError(t, err, "Open success")
	ass.NoError(t, errC, "Close success")
}
