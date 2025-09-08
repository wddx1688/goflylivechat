// Log the panic under unix to the log file

//go:build linux
// +build linux

package tools

import (
	"log"
	"os"

	"golang.org/x/sys/unix"
)

// redirectStderr to the file passed in
func RedirectStderr(f *os.File) {
	err := unix.Dup3(int(f.Fd()), int(os.Stderr.Fd()), 0)
	if err != nil {
		log.Printf("Failed to redirect stderr to file: %v", err)
	}
}
