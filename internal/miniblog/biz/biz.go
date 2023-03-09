package biz

import "github.com/sjxiang/miniblog/internal/miniblog/store"


type Biz interface {

}


type BizImpl struct {
	ds store.IStore
}


// repo db 操作

// restHandler api 校验参数

// dto model 