package cache

import (
	"container/list"
	"strconv"
	"sync"
	"time"

	"github.com/realaryann/keystore/resp"
)

type Cache struct {
	Mut   sync.RWMutex
	Data  map[string][]string
	HData map[string]map[string][]string
	LData map[string]*list.List
}

func (c *Cache) LPop(v resp.Value, mode string) (string, []resp.Value) {
	c.Mut.Lock()
	defer c.Mut.Unlock()
	var bstring string
	barr := []resp.Value{}
	_ = mode
	return bstring, barr

}

func (c *Cache) RPush(v resp.Value) int {
	c.Mut.Lock()
	defer c.Mut.Unlock()
	lname := (v.Array[1]).Bulk
	val := (v.Array[2]).Bulk
	_, exists := c.LData[lname]
	if !exists {
		c.LData[lname] = list.New()
	}
	c.LData[lname].PushBack(val)
	return c.LData[lname].Len()
}

func (c *Cache) LPush(v resp.Value) int {
	c.Mut.Lock()
	defer c.Mut.Unlock()
	lname := (v.Array[1]).Bulk
	val := (v.Array[2]).Bulk
	_, exists := c.LData[lname]
	if !exists {
		c.LData[lname] = list.New()
	}
	c.LData[lname].PushFront(val)
	return c.LData[lname].Len()
}

func (c *Cache) ExpireSet(v resp.Value) int {
	c.Mut.Lock()
	defer c.Mut.Unlock()
	key := (v.Array[1]).Bulk
	if _, ok := c.Data[key]; ok {
		timeslice, _ := strconv.Atoi((v.Array[2]).Bulk)
		exp := int(time.Now().Unix()) + timeslice
		c.Data[key] = append(c.Data[key], strconv.Itoa(exp))
		return 1
	}
	return 0
}

func (c *Cache) HAdd(v resp.Value) int {
	c.Mut.Lock()
	defer c.Mut.Unlock()
	key := (v.Array[1]).Bulk
	// Check if a key already exists, error.
	if _, ok := c.Data[key]; ok {
		return -1
	}
	// HSET key field value.. field value.. field value
	if c.HData[key] == nil {
		c.HData[key] = make(map[string][]string)
	}
	cnt := 0
	for i := 2; i < len(v.Array)-1; i += 2 {
		field := (v.Array[i]).Bulk
		val := (v.Array[i+1]).Bulk
		c.HData[key][field] = append(c.HData[key][field], val)
		cnt++
	}
	return cnt

}

func (c *Cache) HGet(v resp.Value) string {
	// HGET key field
	c.Mut.RLock()
	defer c.Mut.RUnlock()
	key := (v.Array[1]).Bulk
	field := (v.Array[2]).Bulk
	if c.HData[key] == nil {
		return "ERROR"
	}
	return c.HData[key][field][0]
}

func (c *Cache) Incr(v resp.Value) int {
	c.Mut.Lock()
	defer c.Mut.Unlock()
	key := (v.Array[1]).Bulk
	// If Key NOT in cache, create it and add 1.
	val, ok := c.Data[key]
	newv := 1
	if ok {
		tmp, _ := strconv.Atoi(val[0])
		delete(c.Data, key)
		newv = tmp + 1
	}
	c.Data[key] = append(c.Data[key], strconv.Itoa(newv))
	c.Data[key] = append(c.Data[key], strconv.FormatInt((time.Now().Unix()), 10))
	return newv
}

func (c *Cache) Decr(v resp.Value) int {
	c.Mut.Lock()
	defer c.Mut.Unlock()
	key := (v.Array[1]).Bulk
	// If Key NOT in cache, create it and add 1.
	val, ok := c.Data[key]
	newv := -1
	if ok {
		tmp, _ := strconv.Atoi(val[0])
		delete(c.Data, key)
		newv = tmp - 1
	}
	c.Data[key] = append(c.Data[key], strconv.Itoa(newv))
	c.Data[key] = append(c.Data[key], strconv.FormatInt((time.Now().Unix()), 10))
	return newv
}

func (c *Cache) IncrBy(v resp.Value) int {
	c.Mut.Lock()
	defer c.Mut.Unlock()
	key := (v.Array[1]).Bulk
	inc, _ := strconv.Atoi((v.Array[2]).Bulk)
	// If Key NOT in cache, create it and add 1.
	val, ok := c.Data[key]
	newv := inc
	if ok {
		tmp, _ := strconv.Atoi(val[0])
		delete(c.Data, key)
		newv = tmp + inc
	}
	c.Data[key] = append(c.Data[key], strconv.Itoa(newv))
	c.Data[key] = append(c.Data[key], strconv.FormatInt((time.Now().Unix()), 10))
	return newv
}

func (c *Cache) DecrBy(v resp.Value) int {
	c.Mut.Lock()
	defer c.Mut.Unlock()
	key := (v.Array[1]).Bulk
	dec, _ := strconv.Atoi((v.Array[2]).Bulk)
	// If Key NOT in cache, create it and add 1.
	val, ok := c.Data[key]
	newv := dec
	if ok {
		tmp, _ := strconv.Atoi(val[0])
		delete(c.Data, key)
		newv = tmp - dec
	}
	c.Data[key] = append(c.Data[key], strconv.Itoa(newv))
	c.Data[key] = append(c.Data[key], strconv.FormatInt((time.Now().Unix()), 10))
	return newv
}

func (c *Cache) Get(v resp.Value) string {
	c.Mut.RLock()
	defer c.Mut.RUnlock()
	if !c.IsAlive(v) {
		c.Mut.RUnlock()
		c.Del(v)
		c.Mut.RLock()
	}
	val, ok := c.Data[(v.Array[1]).Bulk]
	// Need to add TTL functionality
	if ok {
		return val[0]
	} else {
		return ""
	}
}

func (c *Cache) Add(v resp.Value) {
	c.Mut.Lock()
	defer c.Mut.Unlock()
	key := (v.Array[1]).Bulk
	val := (v.Array[2]).Bulk
	_, ok := c.Data[key]
	if ok {
		delete(c.Data, key)
	}
	// Value, TS
	c.Data[key] = append(c.Data[key], val)
	c.Data[key] = append(c.Data[key], strconv.FormatInt((time.Now().Unix()), 10))
}

func (c *Cache) Del(v resp.Value) int {
	c.Mut.Lock()
	defer c.Mut.Unlock()
	cnt := 0
	for i := range v.Array {
		_, ok := c.Data[v.Array[i].Bulk]
		if ok {
			delete(c.Data, v.Array[i].Bulk)
			cnt++
		}
	}
	return cnt
}

func (c *Cache) IsAlive(v resp.Value) bool {
	// SET, GET, EXISTS has the key as the data[1] value
	c.Mut.RLock()
	defer c.Mut.RUnlock()
	key := (v.Array[1]).Bulk
	if len(c.Data[key]) >= 3 {
		intexpiry, _ := strconv.Atoi(c.Data[key][2])
		if int(time.Now().Unix()) > intexpiry {
			return false
		}
	}
	return true
}
