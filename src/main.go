package main

import (
	"log"
	"net"
)

func main() {
	var status string
	_, err := net.Dial("tcp", "asm.lkimmel.com:20695")
	if err != nil {
		status = "offline"
	} else {
		status = "online"
	}
	log.Println(status)
}
