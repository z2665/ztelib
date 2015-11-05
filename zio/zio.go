package zio

import (
	"errors"
	"io"
	"sync"
)

var ErrShortWrite = errors.New("short write")
var p *sync.Pool

func init() {
	p = &sync.Pool{}
	p.New = func() interface{} {
		return make([]byte, 32*1024)
	}
}

func Copy(dst io.Writer, src io.Reader) (written int64, err error) {

	buf := p.Get().([]byte)
	defer p.Put(buf)
	return copyBuffer(dst, src, buf)
}
func copyBuffer(dst io.Writer, src io.Reader, buf []byte) (written int64, err error) {
	// If the reader has a WriteTo method, use it to do the copy.
	// Avoids an allocation and a copy.
	if wt, ok := src.(io.WriterTo); ok {
		return wt.WriteTo(dst)
	}
	// Similarly, if the writer has a ReadFrom method, use it to do the copy.
	if rt, ok := dst.(io.ReaderFrom); ok {
		return rt.ReadFrom(src)
	}
	if buf == nil {
		buf = make([]byte, 32*1024)
	}
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = ErrShortWrite
				break
			}
		}
		if er == io.EOF {
			break
		}
		if er != nil {
			err = er
			break
		}
	}
	return written, err
}
