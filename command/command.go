package command

import ( 
	"github.com/realaryann/keystore/resp"
	"github.com/realaryann/keystore/cache"
	"strings"
	"strconv"
)

func HandlePing() resp.Value {
	ret := resp.Value{}
	ret.Typ = "string"
	ret.Str = "PONG"
	return ret
}

func HandleInit() resp.Value {
	ret := resp.Value{}
	return ret
}

func HandleExists(v resp.Value, c *cache.Cache) resp.Value {
	// EXISTS arrayelements...
	// Use typ int instead and modify serialize/deserialize
	ret := resp.Value{}
	ret.Typ = "string"
	var cnt int = 1
	for _, val := range(v.Array) {
		tkey := val.Bulk
		_, exists := c.Data[tkey]
		if exists {
			cnt++;
		}
	}
	ret.Str = strconv.Itoa(cnt-1)
	return ret
}

func HandleGet(v resp.Value, c *cache.Cache) resp.Value {
	ret := resp.Value{}
	ret.Typ = "string"
	
	val, ok := c.Data[(v.Array[1]).Bulk]
	// Need to add TTL functionality
	if ok {
		ret.Str = val[0]
	}
	return ret
}

func HandleSet(v resp.Value, c *cache.Cache) resp.Value {
	ret := resp.Value{}
	ret.Typ = "string"
	ret.Str = "+OK"

	c.Add(v)

	return ret
}

func Process(v resp.Value, c *cache.Cache) resp.Value {
	ret := resp.Value{}
	for _, ival := range v.Array {
		if strings.ToUpper(ival.Bulk) == "PING" {
			return HandlePing()
		} else if strings.ToUpper(ival.Bulk) == "DOCS" {
			return HandleInit()
		} else if strings.ToUpper(ival.Bulk) == "SET" {
			return HandleSet(v, c)
		} else if strings.ToUpper(ival.Bulk) == "GET" {
			return HandleGet(v, c)
		} else if strings.ToUpper(ival.Bulk) == "EXISTS" {
			return HandleExists(v, c)
		} else {
			break
		}
	}
	return ret
}