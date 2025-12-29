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

type Writer struct {
	writer io.Writer
}

type Resp struct {
	reader *bufio.Reader
}

func (w *Writer) Write(v Value) {
	var bytes := v.Serialize()

	_, err := w.writer.Write(bytes)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func NewWriter(w io.Writer) *Writer {
	w := Writer{writer: w}
	return &w
}

func NewReader(rd io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(rd)}
}

func (v Value) Serialize() []byte {
	switch (v.typ) {
	case "array":
		return v.SerializeArr()
	case "string":
		return v.SerializeStr()
	case "bulk":
		return v.SerializeBulk()
	case "null":
		return v.SerializeNull()
	case "error":
		return v.SerializeErr()
	default:
		return []byte{}
	}
}

func (v Value) SerializeArr() []byte {
	var bytes []byte
	length := len(v.array)
	bytes = append(bytes, ARRAY)
	bytes = append(bytes, []byte(strconv.Itoa(length))...)
	bytes = append(bytes, '\r', '\n')
	
	for i := 0; i<length; i++ {
		bytes = append(bytes, []byte(v.array[i].Serialize())...)
	}
	return bytes
}

func (v Value) SerializeErr() []byte {
	var bytes []byte
	bytes = append(bytes, ERROR)
	bytes = append(bytes, []byte(v.str)...)
	bytes = append(bytes, '\r', '\n')
	return bytes
}

func (v Value) SerializeNull() []byte {
	return []byte("$-1\r\n")
}

func (v Value) SerializeStr() []byte {
	var bytes []byte
	bytes = append(bytes, STRING)
	bytes = append(bytes, []byte(v.str)...)
	bytes = append(bytes, '\r','\n')
	return bytes
}

func (v Value) SerializeBulk() []byte {
	var bytes []byte
	bytes = append(bytes, BULK)
	bytes = append(bytes, []byte(strconv.Itoa(len(v.bulk)))...)
	bytes = append(bytes, '\r', '\n')
	bytes = append(bytes, []byte(v.bulk)...)
	bytes = append(bytes, '\r', '\n')
	return bytes
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
	// EG: *2\r\n$5\r\nhello\r\n$5\r\nworld\r\n
	v := Value{}
	v.typ = "array"

	length, _ := r.ReadInteger()
	v.array = make([]Value, length)
	
	for i := 0; i<length; i++ {
		// Read the resulting parts of the array
		val, err := r.Read()
		if err != nil {
			fmt.Println("Error: err")
			return v, nil
		}
		v.array[i] = val
	}
	return v, nil
}

func (r* Resp) ReadBulk() (Value, error) {
	// TODO
	// Eg: $5\r\nhello\r\n
	var v Value
	v.typ = "bulk"

	length, _ := r.ReadInteger()

	bulk := make([]byte, length)

	r.reader.Read(bulk)

	v.bulk = string(bulk)

	r.ReadLine()
	
	return v, nil
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