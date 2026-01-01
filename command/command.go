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
		}
	}
	return ret
}