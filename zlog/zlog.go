package zlog

import (
	_ "fmt"
	"io"
	"sync"
	"sync/atomic"
)

type Zlog struct {
	mutex   sync.Mutex
	bufsize int64
	buf     []byte
	wt      io.Writer
}

func New(out io.Writer, buffersize int64) *Zlog {
	return &Zlog{wt: out, bufsize: buffersize, buf: make([]byte, buffersize)}
}

//重新分配缓冲区的大小
func (this *Zlog) Resize(newsize int64) {
	atomic.StoreInt64(&this.bufsize, newsize)
	this.buf = make([]byte, this.bufsize)
}
func (this *Zlog) GetBufSize() int64 {
	return this.bufsize
}
