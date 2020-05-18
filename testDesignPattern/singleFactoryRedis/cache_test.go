package singleFactoryRedis

import (
	"fmt"
	"testing"
)

func TestCacheFactory_Create(t *testing.T) {
	cacheFactory := &CacheFactory{}
	c1, err := cacheFactory.Create(redis)
	if err != nil {
		fmt.Println(err.Error())
	}
	c1.Set("a", "a")
	c1.Get("a")

	c2, err := cacheFactory.Create(mem)
	if err != nil {
		fmt.Println(err.Error())
	}
	c2.Set("n", "aaaaaaaa")
	c2.Get("n")

}
