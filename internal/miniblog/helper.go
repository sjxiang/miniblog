package miniblog

import (
	"os"
	"syscall"

	"github.com/sjxiang/miniblog/internal/pkg/log"
)

func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     false,
		DisableStacktrace: false,
		Level:             env.LogLevel,
		Format:            env.LogFormat,
		OutputPaths:       env.LogOutputPaths,
	}
}

// 定义要监听的目标信号
var signals = []os.Signal{
					  // kill -9，发送 syscall.SIGKILL 信号，但不能被捕获
	syscall.SIGINT,   // kill -2，即常用的 CTRL + c，发送 syscall.SIGINT 信号
	syscall.SIGTERM,  // kill 默认会发送 syscall.SIGTERM 信号 
}
