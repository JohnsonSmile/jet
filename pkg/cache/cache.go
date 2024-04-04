package cache

import "time"

type Cache interface {
	// size: 1KB 100KB 1MB 2MB 1GB
	SetMaxMemory(size Size) error
	// 将value写入缓存
	Set(key string, val any) error
	// 将value写入缓存
	SetEx(key string, val any, expire time.Duration) error
	// 根据key获取value
	Get(key string) (val any, err error)
	// 删除key
	Del(key string) error
	// 情况所有key
	Flush() error
	// 获取缓存的所有的key的数量
	Count() int64
}
