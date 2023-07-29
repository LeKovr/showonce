package main_test

import (
	"context"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/LeKovr/go-kit/config"
	cmd "github.com/LeKovr/showonce/cmd/showonce"
)

func TestRunErrors(t *testing.T) {
	// Save original args
	a := os.Args
	_, p2 := GetPorts(t)

	tests := []struct {
		name string
		code int
		args []string
	}{
		{"Help", config.ExitHelp, []string{"-h"}},
		{"UnknownFlag", config.ExitBadArgs, []string{"-0"}},
		{"IncorrectEndPoint", config.ExitError, []string{"--log.debug",
			"--listen", "xx:unknown",
			"--listen_grpc", fmt.Sprintf(":%d", p2),
		}},
	}
	ctx := context.Background()
	for _, tt := range tests {
		os.Args = append([]string{a[0]}, tt.args...)

		var c int

		cmd.Run(ctx, func(code int) { c = code })
		assert.Equal(t, tt.code, c, tt.name)
	}

	// Restore original args
	os.Args = a
}

func TestRun(t *testing.T) {
	p1, p2 := GetPorts(t)
	// Save original args
	a := os.Args
	os.Args = append([]string{a[0]},
		"--listen", fmt.Sprintf(":%d", p1),
		"--listen_grpc", fmt.Sprintf(":%d", p2),
	)
	var c int
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	cmd.Run(ctx, func(code int) { c = code })
	assert.Equal(t, config.ExitNormal, c, "Normal run")
	// Restore original args
	os.Args = a
}

func GetPorts(t *testing.T) (int, int) {
	// Find ports
	p1, err := GetFreePort()
	assert.NoError(t, err, "Port")
	p2, err := GetFreePort()
	assert.NoError(t, err, "Port2")
	return p1, p2
}

// Code from https://gist.github.com/sevkin/96bdae9274465b2d09191384f86ef39d

// GetFreePort asks the kernel for a free open port that is ready to use.
func GetFreePort() (int, error) {
	var a *net.TCPAddr
	var err error
	if a, err = net.ResolveTCPAddr("tcp", "localhost:0"); err == nil {
		var l *net.TCPListener
		if l, err = net.ListenTCP("tcp", a); err == nil {
			port := l.Addr().(*net.TCPAddr).Port
			l.Close()
			return port, err
		}
	}
	return 0, err
}
