package zlog

import (
	"io"
	"fmt"
	"log"
)
type Zlog struct{
	var ln *log.Logger//notic日志
	var lw *log.Logger//警告日志
	var le *log.Logger//错误日志
}

func New(out io.Writer){
	return &Zlog{ln: log.New(out,"notic",)}
}