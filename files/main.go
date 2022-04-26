package main

import (
	"fmt"
	"os"
)

func main() {
	// Files
}
func killServer(pidFile string) error {
	// os.Open - read
	// os.Write - write & truncate
	// os.OpenFile - append (and more)
	// Go idiom: acquire resource, check for error, defer release
	file, err := os.Open(pidFile)
	if err != nil {
		return err
	}
	// defer handles resources, file.Close() will be called when we exit killServer
	defer file.Close()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return err
	}
}
