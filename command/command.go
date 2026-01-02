package command

import ( 
	"github.com/realaryann/keystore/resp"
	"time"
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

func HandleGet(v resp.Value, cache map[string][]string) resp.Value {
	ret := resp.Value{}
	ret.Typ = "string"
	
	val, ok := cache[(v.Array[1]).Bulk]
	// Need to add TTL functionality
	if ok {
		ret.Str = val[0]
	} else {
		ret.Str = "ERROR"
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
		if ival.Bulk == "PING" {
			return HandlePing()
		} else if ival.Bulk == "DOCS" {
			return HandleInit()
		} else if ival.Bulk == "SET" {
			return HandleSet(v, cache)
		} else if ival.Bulk == "GET" {
			return HandleGet(v, cache)
		}
	}
	return ret
}