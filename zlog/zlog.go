package zlog

import (
	"fmt"
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

func (this *Zlog) Notice(v ...interface{}) {
	s := fmt.Sprint(v...)
	this.logdo(s, "Notice: ")
}
func (this *Zlog) logdo(s string, title string) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.buf = this.buf[:0] //清空
	this.buf = append(this.buf, title...)
	this.buf = append(this.buf, s...)
	this.wt.Write(this.buf)
}
