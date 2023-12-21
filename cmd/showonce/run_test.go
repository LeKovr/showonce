package main_test

import (
	"context"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	ass "github.com/alecthomas/assert/v2"

	"github.com/LeKovr/go-kit/config"
	cmd "github.com/LeKovr/showonce/cmd/showonce"
)

func TestRunErrors(t *testing.T) {
	// Save original args
	a := os.Args
	ports := GetPorts(t)

	tests := []struct {
		name string
		code int
		args []string
	}{
		{"Help", config.ExitHelp, []string{"-h"}},
		{"UnknownFlag", config.ExitBadArgs, []string{"-0"}},
		{"IncorrectEndPoint", config.ExitError, []string{
			"--log.debug",
			"--listen", "xx:unknown",
			"--listen_grpc", fmt.Sprintf(":%d", ports[0]),
		}},
	}
	ctx := context.Background()
	for _, tt := range tests {
		os.Args = append([]string{a[0]}, tt.args...)
		var c int
		cmd.Run(ctx, func(code int) { c = code })
		ass.Equal(t, tt.code, c, tt.name)
	}

	// Restore original args
	os.Args = a
}

func TestRun(t *testing.T) {
	ports := GetPorts(t)
	// Save original args
	a := os.Args
	os.Args = append([]string{a[0]},
		"--listen", fmt.Sprintf(":%d", ports[0]),
		"--listen_grpc", fmt.Sprintf(":%d", ports[1]),
	)
	var c int
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	cmd.Run(ctx, func(code int) { c = code })
	ass.Equal(t, config.ExitNormal, c, "Normal run")
	// Restore original args
	os.Args = a
}

func GetPorts(t *testing.T) []int {
	t.Helper()
	// Find ports
	p1, err := GetFreePort()
	ass.NoError(t, err, "Port")
	p2, err := GetFreePort()
	ass.NoError(t, err, "Port2")
	return []int{p1, p2}
}

// Code from https://gist.github.com/sevkin/96bdae9274465b2d09191384f86ef39d

// GetFreePort asks the kernel for a free open port that is ready to use.
func GetFreePort() (int, error) {
	var a *net.TCPAddr
	var err error
	if a, err = net.ResolveTCPAddr("tcp", "localhost:0"); err == nil {
		var l *net.TCPListener
		if l, err = net.ListenTCP("tcp", a); err == nil {
			// sample for https://www.joeshaw.org/dont-defer-close-on-writable-files/
			if port, ok := l.Addr().(*net.TCPAddr); ok {
				err = l.Close()
				return port.Port, err
			}
			err = l.Close()
		}
	}
	return 0, err
}
