package singleFactoryRedis

import (
	"errors"
	"fmt"
)

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

type CacheFactory struct {
}

const (
	redis = iota
	mem
)

type CacheType int

func (c *CacheFactory) Create(cacheType CacheType) (Cache, error) {
	if cacheType == redis {
		return &Redis{data: map[string]string{}}, nil
	}
	if cacheType == mem {
		return &MemCache{data: map[string]string{}}, nil
	}
	return nil, errors.New("Cache Type not support.")
}
