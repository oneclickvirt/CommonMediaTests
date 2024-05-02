package main

import (
	"os"
	"testing"
)

func TestMain_en(t *testing.T) {
	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
	}()
	os.Args = []string{"", "-l", "en"}
	main()
}

func TestMain_zh(t *testing.T) {
	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
	}()
	os.Args = []string{"", "-l", "zh"}
	main()
}
