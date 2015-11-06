package zlog

import (
	"bytes"
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
func TestNotice(t *testing.T) {
	var buf bytes.Buffer
	log := New(&buf, 200)
	log.Notice("hello")
	if buf.String() != "Notice: hello" {
		t.Errorf("Err String %s", buf.String())
	}
}
