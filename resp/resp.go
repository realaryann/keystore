package resp

import "fmt"

const (
	STRING  = '+'
	ERROR   = '-'
	INTEGER = ':'
	BULK    = '$'
	ARRAY   = '*'
)

func Resp() {
	fmt.Println("Resp Serializer/Deserializer")

}
