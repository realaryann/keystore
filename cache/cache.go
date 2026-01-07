package cache

import (
	"github.com/realaryann/keystore/resp"
	"sync"
	"strconv"
	"time"
	"fmt"
)

type Cache struct {
	Mut sync.Mutex
	Data map[string][]string
}


func (c* Cache) ExpireSet(v resp.Value) {
	/*
	c.Mut.Lock()
	key := (v.Array[1]).Bulk
	timeslice, _ := strconv.Atoi((v.Array[2]).Bulk)
	orig := c.Data[key][1]
	orig = c.Data[key][1].Add(timeslice*time.Second)
	fmt.Println(c.Data[key][1])
	fmt.Println(orig)
	c.Data[key][2] = orig
	c.Mut.Unlock()
	*/
}

func (c *Cache) Add(v resp.Value) {
	c.Mut.Lock()
	key := (v.Array[1]).Bulk
	val := (v.Array[2]).Bulk
	// Value, TS
	c.Data[key] = append(c.Data[key], val)
	c.Data[key] = append(c.Data[key], strconv.FormatInt((time.Now().Unix()), 10))
	fmt.Println(strconv.FormatInt((time.Now().Unix()), 10))
	c.Mut.Unlock()
}

func (c * Cache) Del(v resp.Value)  {
	c.Mut.Lock()
	for i := range(v.Array) {
		delete(c.Data, v.Array[i].Bulk)
	}
	c.Mut.Unlock()
}
