package main

import (
	"os"

	"github.com/sjxiang/miniblog/internal/miniblog"
)

// Go 程序的默认入口主函数
func main() {
	command := miniblog.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
