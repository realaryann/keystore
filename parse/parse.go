package parse

import (
	"flag"
)

func Parse() *string {
	tcp_port := flag.String("p", "6000", "Custom port number for KeyStore. Default: 6000")
	flag.Parse()
	return tcp_port
}