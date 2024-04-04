package memory

import (
	"github.com/JohnsonSmile/jet/pkg/cache"
	"time"
)

type Cache struct {
	maxSize cache.Size
	curSize cache.Size
}

func NewMemoryCache() *Cache {
	return &Cache{}
}

func (m *Cache) SetMaxMemory(size cache.Size) error {
	if size < 0 {
		return ErrInvalidMaxSize
	}
	m.maxSize = size
	return nil
}

func (m *Cache) Set(key string, val any) error {
	//TODO implement me
	panic("implement me")
}

func (m *Cache) SetEx(key string, val any, expire time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (m *Cache) Get(key string) (val any, err error) {
	//TODO implement me
	panic("implement me")
}

func (m *Cache) Del(key string) error {
	//TODO implement me
	panic("implement me")
}

func (m *Cache) Flush() error {
	//TODO implement me
	panic("implement me")
}

func (m *Cache) Count() int64 {
	//TODO implement me
	panic("implement me")
}

var _ cache.Cache = (*Cache)(nil)
