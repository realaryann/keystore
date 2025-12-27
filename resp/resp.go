package resp

import (
	"io"
	"bufio"
	"fmt"
	"strconv"
)

const (
	STRING  = '+'
	ERROR   = '-'
	INTEGER = ':'
	BULK    = '$'
	ARRAY   = '*'
)

type Value struct {
	typ   string
	str   string
	num   int
	bulk  string
	array []Value
}

type Resp struct {
	reader *bufio.Reader
}

func NewReader(rd io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(rd)}
}

func (r *Resp) Read() (Value, error) {
	sym, err := r.reader.ReadByte()
	if err != nil {
		fmt.Println("Error: ", err)
		return Value{}, err
	}
	switch sym {
	case ARRAY:
		return r.ReadArray()
	case BULK:
		return r.ReadBulk()
	default:
		fmt.Println("Unknown Type Symbol: ", sym)
		return Value{}, nil
	}

}

func (r *Resp) ReadArray() (Value, error) {
	// TODO
	return Value{}, nil
}

func (r* Resp) ReadBulk() (Value, error) {
	// TODO
	return Value{}, nil
}

func (r *Resp) ReadLine() ([]byte, error) {
	// Go thru the buffer, "$5\r\nAryan\r\n"
	var line []byte
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			fmt.Println("Error: ", err)
			return line, err
		}
		line = append(line, b)
		if len(line) >= 2 && line[len(line)-2] == '\r' {
			break
		}
	}
	return line[:len(line)-2], nil
}

func (r *Resp) ReadInteger() (int, error) {
	line, err := r.ReadLine()
	if err != nil {
		fmt.Println("Error: ", err)
		return 0, err
	}
	// strconv.ParseInt (string, base of string, return type bit size)
	i, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		fmt.Println("Error: ", err)
		return 0, err
	}
	return int(i), nil
}