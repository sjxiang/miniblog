package log


import (
	"sync"

	"go.uber.org/zap"
)

type zapLogger struct {
	z *zap.Logger
}

type Options struct{

}

var (
	mu sync.Mutex
	
)


func Init()