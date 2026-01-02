package command

import ( 
	"github.com/realaryann/keystore/resp"
	"time"
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

func HandleExists(v resp.Value, cache map[string][]string) resp.Value {
	// EXISTS arrayelements...
	// Use typ int instead and modify serialize/deserialize
	ret := resp.Value{}
	ret.Typ = "string"
	var cnt int = 1
	for _, val := range(v.Array) {
		tkey := val.Bulk
		_, exists := cache[tkey]
		if exists {
			cnt++;
		}
	}
	ret.Str = strconv.Itoa(cnt-1)
	return ret
}

func HandleGet(v resp.Value, cache map[string][]string) resp.Value {
	ret := resp.Value{}
	ret.Typ = "string"
	
	val, ok := cache[(v.Array[1]).Bulk]
	// Need to add TTL functionality
	if ok {
		ret.Str = val[0]
	}
	return ret
}

func HandleSet(v resp.Value, cache map[string][]string) resp.Value {
	ret := resp.Value{}
	ret.Typ = "string"
	ret.Str = "+OK"

	key := (v.Array[1]).Bulk
	val := (v.Array[2]).Bulk
	// Value, TS
	cache[key] = append(cache[key], val)
	cache[key] = append(cache[key], (time.Now()).Format(time.TimeOnly))

	return ret
}

func Process(v resp.Value, cache map[string][]string) resp.Value {
	ret := resp.Value{}
	for _, ival := range v.Array {
		if strings.ToUpper(ival.Bulk) == "PING" {
			return HandlePing()
		} else if strings.ToUpper(ival.Bulk) == "DOCS" {
			return HandleInit()
		} else if strings.ToUpper(ival.Bulk) == "SET" {
			return HandleSet(v, cache)
		} else if strings.ToUpper(ival.Bulk) == "GET" {
			return HandleGet(v, cache)
		} else if strings.ToUpper(ival.Bulk) == "EXISTS" {
			return HandleExists(v, cache)
		}
	}
	return ret
}