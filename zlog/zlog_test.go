package zlog

import (
	"os"
	"testing"
)

func TestResize(t *testing.T) {
	log := New(os.Stdout, 200)
	log.Resize(400)
	if log.bufsize != 400 {
		t.Error("Resize err")
	}
}
