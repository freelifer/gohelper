package cache

import (
	"time"
)

// get cached value by key.
func Get(key string) interface{} {
	if C != nil {
		return C.Get(key)
	}
	return nil
}

// set cached value with key and expire time.
func Put(key string, val interface{}, timeout time.Duration) error {
	if C != nil {
		return C.Put(key, val, timeout)
	}
	return nil
}

// delete cached value by key.
func Delete(key string) error {
	if C != nil {
		return C.Delete(key)
	}
	return nil
}

// check if cached value exists or not.
func IsExist(key string) bool {
	if C != nil {
		return C.IsExist(key)
	}
	return false
}
