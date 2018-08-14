// +build redis

package cache

import (
	"github.com/freelifer/cache"
	_ "github.com/freelifer/cache/redis"
)

var C cache.Cache

func init() {
	var err error
	C, err = cache.NewCache("redis", `{"key":"collectionName","conn":":6039","dbNum":"0","password":"thePassWord"}`)
	if err != nil {
		panic(err)
	}
}
