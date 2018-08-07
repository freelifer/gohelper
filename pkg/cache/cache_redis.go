// +build redis

package cache

import (
	"github.com/freelifer/cache"
	_ "github.com/garyburd/redigo"
)

var c cache.Cache

func init() {
	c, err := cache.NewCache("redis", `{"key":"collectionName","conn":":6039","dbNum":"0","password":"thePassWord"}`)
	if err != nil {
		panic(err)
	}
}
