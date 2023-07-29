//go:build !test
// +build !test

// This file holds code which does not covered by tests

package main

import (
	"context"
	"os"
)

func main() {
	Run(context.Background(), os.Exit)
}
