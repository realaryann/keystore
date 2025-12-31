package command

import ( 
	"github.com/realaryann/keystore/resp"
)


func Process(v *resp.Value) {
	for _, ival := range v.Array {
		if ival.Bulk == "PING" {
			v.Typ = "string"
			v.Str = "PONG"
			v.Array = []resp.Value{}
			break
		} else if ival.Bulk == "DOCS" {
			v.Array = []resp.Value{}
			break
		}
	}

}