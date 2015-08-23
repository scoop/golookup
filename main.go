package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 1 || os.Args[1] == "" {
		log.Fatalln("Please provide a DNS name as the first argument.")
	}

	host := os.Args[1]
	addrs, err := net.LookupHost(host)

	if err != nil {
		log.Fatalf("Error resolving host: %v", err)
	}

	fmt.Println(strings.Join(addrs, ","))
}
