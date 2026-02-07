package command

import ( 
	"github.com/realaryann/keystore/resp"
	"github.com/realaryann/keystore/cache"
	"strings"
)

func HandlePing() resp.Value {
	ret := resp.Value{}
	ret.Typ = "string"
	ret.Str = "PONG"
	return ret
}

func HandleInit() resp.Value {
	ret := resp.Value{}
	ret.Typ = "string"
	ret.Str = "OK"
	return ret
}

func HandleExists(v resp.Value, c *cache.Cache) resp.Value {
	// EXISTS arrayelements...
	// Use typ int instead and modify serialize/deserialize
	ret := resp.Value{}
	ret.Typ = "integer"
	var cnt int = 1
	for _, val := range(v.Array) {
		tkey := val.Bulk
		_, exists := c.Data[tkey]
		if exists {
			cnt++;
		}
	}
	ret.Num = cnt-1
	return ret
}

func HandleGet(v resp.Value, c *cache.Cache) resp.Value {
	ret := resp.Value{}
	ret.Typ = "bulk"
	ret.Bulk = c.Get(v)
	return ret
}

func HandleSet(v resp.Value, c *cache.Cache) resp.Value {
	ret := resp.Value{}
	ret.Typ = "string"
	ret.Str = "OK"
	
	c.Add(v)

	return ret
}

func HandleDel(v resp.Value, c *cache.Cache) resp.Value {
	ret := resp.Value{}
	ret.Typ = "integer"

	ret.Num = c.Del(v)

	return ret
}

func HandleExpire(v resp.Value, c *cache.Cache) resp.Value {
	ret := resp.Value{}
	ret.Typ = "integer"
	ret.Num = c.ExpireSet(v)
	return ret
}

func HandleBench() resp.Value {
	ret := resp.Value{}
	ret.Typ = "array"
	return ret
}

func HandleHset(v resp.Value, c *cache.Cache) resp.Value {
	ret := resp.Value{}
	ret.Typ = "integer"
	ret.Num = c.HAdd(v)
	return ret
}

func HandleHget(v resp.Value, c *cache.Cache) resp.Value {
	ret := resp.Value{}
	ret.Typ = "bulk"
	ret.Bulk = c.HGet(v)
	return ret
}

func HandleIncr(v resp.Value, c *cache.Cache) resp.Value {
	ret := resp.Value{}
	ret.Typ = "integer"
	ret.Num = c.Incr(v)
	return ret
}

func HandleDecr(v resp.Value, c *cache.Cache) resp.Value {
	ret := resp.Value{}
	ret.Typ = "integer"
	ret.Num = c.Decr(v)
	return ret
}

func Process(v resp.Value, c *cache.Cache) resp.Value {
	ret := resp.Value{}
	for _, ival := range v.Array {
		switch(strings.ToUpper(ival.Bulk)) {
		case "PING":
			return HandlePing()
		case "COMMAND":
			return HandleInit()
		case "HELLO":
			return HandleInit()
		case "CONFIG":
			return HandleBench()
		case "SET":
			return HandleSet(v, c)
		case "GET":
			return HandleGet(v, c)
		case "EXISTS":
			return HandleExists(v, c)
		case "DEL":
			return HandleDel(v, c)
		case "EXPIRE":
			return HandleExpire(v, c)
		case "HSET":
			return HandleHset(v, c)
		case "HGET":
			return HandleHget(v, c)
		case "INCR":
			return HandleIncr(v, c)
		case "DECR":
			return HandleDecr(v, c)
		default:
			break
		}
		
	}
	return ret
}