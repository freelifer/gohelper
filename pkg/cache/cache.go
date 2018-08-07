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

// get cached value by key.
func Get(key string) interface{} {

}
	// GetMulti is a batch version of Get.
func GetMulti(keys []string) []interface{} {

}
	// set cached value with key and expire time.
func Put(key string, val interface{}, timeout time.Duration) error {

}
	// delete cached value by key.
func Delete(key string) error {

}

	// check if cached value exists or not.
	IsExist(key string) bool {

	}
