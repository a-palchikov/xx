package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	origName := fmt.Sprint(os.Args[0], "-orig")
	cmd := []string{"/usr/bin/qemu-aarch64-static", "-L", ".", origName}
	cmd = append(cmd, os.Args[1:]...)
	log.Printf("Starting %v\n", cmd)
	if err := syscall.Exec("/usr/bin/qemu-aarch64-static", cmd, nil); err != nil {
		log.Fatal("Failed to execve: ", err)
	}
}
