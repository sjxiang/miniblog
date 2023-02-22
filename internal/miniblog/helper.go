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
	syscall.SIGKILL, // kill -9 pid，该信号不能被捕获 
	syscall.SIGINT,  // kill -2，等同于 CTRL + c
	syscall.SIGTERM, // kill pid
	syscall.SIGQUIT, // ctrl + \
}

