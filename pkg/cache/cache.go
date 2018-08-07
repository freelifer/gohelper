package cache

import (
	"time"
)

// get cached value by key.
func Get(key string) interface{} {
	if c != nil {
		return c.Get(key)
	}
	return nil
}

// set cached value with key and expire time.
func Put(key string, val interface{}, timeout time.Duration) error {
	if c != nil {
		return c.Put(key, val, timeout)
	}
	return nil
}

// delete cached value by key.
func Delete(key string) error {
	if c != nil {
		return c.Delete(key)
	}
	return nil
}

// check if cached value exists or not.
func IsExist(key string) bool {
	if c != nil {
		return c.IsExist(key)
	}
	return false
}
