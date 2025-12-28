package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"github.com/realaryann/keystore/resp"
)

func main() {
	r := strings.NewReader("$5\r\nhello\r\n")
	obj := resp.NewReader(r)
	obj.Read()
	fmt.Println("KeyStore Server")
	tcpl, err := net.Listen("tcp", ":6000")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	conn, err := tcpl.Accept()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer conn.Close()

	// Infinite Loop to answer connections
	for {
		buffer := make([]byte, 1024)

		_, err = conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error: ", err.Error())
			os.Exit(1)
		}

		conn.Write([]byte("+OK\r\n"))
	}

}
