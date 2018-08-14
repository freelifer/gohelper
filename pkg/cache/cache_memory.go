// +build !redis

package cache

import (
	"github.com/freelifer/cache"
)

var C cache.Cache

func init() {
	var err error
	C, err = cache.NewCache("memory", `{"interval":60}`)

	if err != nil {
		panic(err)
	}
}
