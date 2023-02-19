package miniblog

import "github.com/sjxiang/miniblog/internal/pkg/log"

func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     false,
		DisableStacktrace: false,
		Level:             env.LogLevel,
		Format:            env.LogFormat,
		OutputPaths:       env.LogOutputPaths,
	}
}
