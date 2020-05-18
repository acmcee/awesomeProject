package factory

import "testing"

func TestCacheFactory_Create(t *testing.T) {
	var cacheFactory CacheFactory
	cacheFactory = &RedisCacheFactory{}
	redisCache := cacheFactory.Create()
	redisCache.Set("a", "a")
	redisCache.Get("a")

}
