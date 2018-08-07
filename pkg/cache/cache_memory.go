// +build !redis

package cache

import (
	"github.com/freelifer/cache"
)

var c cache.Cache

func init() {
	c, err := cache.NewCache("memory", `{"interval":60}`)

	if err != nil {
		panic(err)
	}
}
