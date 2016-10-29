package main

import (
	"os"
	"syscall"
)

func main() { //cMove A B
	err := os.Link(os.Args[1], os.Args[2])
	if err != nil {
		return
	}
	syscall.Unlink(os.Args[1])
}
