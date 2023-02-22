package log

import "go.uber.org/zap/zapcore"

// Custom 定制

// 设置 Log Entry 格式
func customEncoder() *zapcore.Encoder {
	return nil
}

// 设置日志输出流
func customLogWriter() zapcore.WriteSyncer {
	return nil
}


