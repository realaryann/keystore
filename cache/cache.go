package cache

import (
	"github.com/realaryann/keystore/resp"
	"sync"
	"strconv"
	"time"
)

type Cache struct {
	Mut sync.Mutex
	Data map[string][]string
}


func (c* Cache) ExpireSet(v resp.Value) {
	c.Mut.Lock()
	key := (v.Array[1]).Bulk
	timeslice, _ := strconv.Atoi((v.Array[2]).Bulk)
	orig, _ := strconv.Atoi(c.Data[key][1])
	exp := orig+timeslice
	c.Data[key] = append(c.Data[key], strconv.Itoa(exp))
	c.Mut.Unlock()
}

func (c *Cache) Add(v resp.Value) {
	c.Mut.Lock()
	key := (v.Array[1]).Bulk
	val := (v.Array[2]).Bulk
	// Value, TS
	c.Data[key] = append(c.Data[key], val)
	c.Data[key] = append(c.Data[key], strconv.FormatInt((time.Now().Unix()), 10))
	c.Mut.Unlock()
}

func (c *Cache) Del(v resp.Value)  {
	c.Mut.Lock()
	for i := range(v.Array) {
		delete(c.Data, v.Array[i].Bulk)
	}
	c.Mut.Unlock()
}

func (c *Cache) IsAlive(v resp.Value) bool {
	// SET, GET, EXISTS has the key as the data[1] value
	key := (v.Array[1]).Bulk
	if len(c.Data[key]) >= 3 {
		intexpiry, _ := strconv.Atoi(c.Data[key][2])
		if int(time.Now().Unix()) > intexpiry {
			return false
		}
	}
	return true
}
