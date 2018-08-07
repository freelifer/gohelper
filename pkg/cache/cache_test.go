package cache

import (
	"github.com/freelifer/cache"
	"testing"
	"time"
)

func Test_Put(t *testing.T) {
	c, err := cache.NewCache("memory", `{"interval":60}`)

	if err != nil {
		t.Error(err)
	}

	c.Put("key", "val", 1*time.Second)

	val := c.Get("key")

	if val != "val" {
		t.Errorf("Error get key val ！= val")
	}
}
