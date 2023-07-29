package static_test

import (
	"testing"

	"github.com/LeKovr/showonce/static"
	"github.com/stretchr/testify/require"
)

func TestNewEmbed(t *testing.T) {
	got, err := static.New("")
	require.NotNil(t, got, "FS not nil")
	require.NoError(t, err, "New success")
	f, err := got.Open("index.html")
	errC := f.Close()
	require.NoError(t, err, "Open success")
	require.NoError(t, errC, "Close success")
}

func TestNewReal(t *testing.T) {
	got, err := static.New("js")
	require.NotNil(t, got, "FS not nil")
	require.NoError(t, err, "New success")
	f, err := got.Open("service.swagger.json")
	errC := f.Close()
	require.NoError(t, err, "Open success")
	require.NoError(t, errC, "Close success")
}
