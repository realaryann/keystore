package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
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
