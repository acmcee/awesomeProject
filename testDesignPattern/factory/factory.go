package factory

import "fmt"

type Cache interface {
	Get(name string)
	Set(key, value string)
}

// redis

type Redis struct {
	data map[string]string
}

func (p *Redis) Get(key string) {
	fmt.Println("redis key:", p.data[key])
}

func (p *Redis) Set(key, value string) {
	p.data[key] = value
}

// MemCache
type MemCache struct {
	data map[string]string
}

func (p *MemCache) Get(key string) {
	fmt.Println("memcache key:", p.data[key])
}

func (p *MemCache) Set(key, value string) {
	p.data[key] = value
}

type CacheFactory interface {
	Create() Cache
}

type RedisCacheFactory struct {
}

func (r *RedisCacheFactory) Create() Cache {
	return &Redis{data: map[string]string{}}
}

type MemCacheFactory struct {
}

func (m *MemCacheFactory) Create() Cache {
	return &MemCache{data: map[string]string{}}
}
