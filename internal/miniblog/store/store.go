package store


import (
	"sync"

	"gorm.io/gorm"
)

var (
	once sync.Once

	// 全局变量，方便其它包直接调用已初始化好的 S 实例
	S *datastore
)


// IStore 定义了 Store 层需要实现的方法
type IStore interface {
	Users() UserStore
}

// 确保 datastore 实现了 IStore 接口
var _ IStore = (*datastore)(nil)

// datastore 是 IStore 的一个具体实现
type datastore struct {
	db *gorm.DB
}

// NewStore 创建一个 IStore 类型的实例
func NewStore(db *gorm.DB) *datastore {
	// 确保 S 只被初始化一次
	once.Do(func() {
		S = &datastore{
			db: db,
		}
	})

	return S
}

// Users 返回了一个实现了 UserStore 接口的实例
func (ds *datastore) Users() UserStore {
	return NewUsers(ds.db)
}


// Store interface
// StoreImpl struct
// func NewStoreImpl
// impl StoreImpl