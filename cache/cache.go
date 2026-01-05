package cache

import (
	"github.com/realaryann/keystore/resp"
	"sync"
	"time"
)

type Cache struct {
	Mut sync.Mutex
	Data map[string][]string
}

func (c *Cache) Add(v resp.Value) {
	c.Mut.Lock()
	key := (v.Array[1]).Bulk
	val := (v.Array[2]).Bulk
	// Value, TS
	c.Data[key] = append(c.Data[key], val)
	c.Data[key] = append(c.Data[key], (time.Now()).Format(time.TimeOnly))
	c.Mut.Unlock()
}

func (c * Cache) Del(v resp.Value)  {
	c.Mut.Lock()
	for i := range(v.Array) {
		delete(c.Data, v.Array[i].Bulk)
	}
	c.Mut.Unlock()
}
