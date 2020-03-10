package main

import (
	"testing"
	"time"
)

func TestRunMain(t *testing.T) {
 go main()
 time.Sleep(100 * time.Millisecond)
}
