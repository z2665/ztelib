package ztlib

import (
	"bytes"
	"io"
	"testing"
)

type Buffer struct {
	bytes.Buffer
	io.ReaderFrom // conflicts with and hides bytes.Buffer's ReaderFrom.
	io.WriterTo   // conflicts with and hides bytes.Buffer's WriterTo.
}

func TestCopy(t *testing.T) {
	rb := new(Buffer)
	wb := new(Buffer)
	rb.WriteString("hello, world.")
	Copy(wb, rb)
	if wb.String() != "hello, world." {
		t.Errorf("Copy did not work properly")
	}
}
func BenchmarkCopy(b *testing.B) {
	rb := new(Buffer)
	wb := new(Buffer)
	rb.WriteString("hello, world.")

	for i := 0; i < b.N; i++ { //use b.N for looping
		Copy(wb, rb)
	}
}
func BenchmarkIOCopy(b *testing.B) {
	rb := new(Buffer)
	wb := new(Buffer)
	rb.WriteString("hello, world.")
	for i := 0; i < b.N; i++ { //use b.N for looping
		io.Copy(wb, rb)
	}
}
