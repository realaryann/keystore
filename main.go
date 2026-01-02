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
	// Key: [value, timestamp, expiry]
	cache := make(map[string][]string)
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
		
		
		retval := command.Process(value, cache)
		
		write := resp.NewWriter(conn)
		write.Write(retval)

	}

}
