package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"github.com/realaryann/keystore/resp"
	"github.com/realaryann/keystore/command"
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
		response := resp.NewReader(conn)	

		value, err := response.Read()

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error: ", err.Error())
			os.Exit(1)
		}
		
		
		command.Process(&value)
		
		write := resp.NewWriter(conn)
		write.Write(value)

	}

}
