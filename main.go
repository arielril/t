package main

import (
	"fmt"
	"os"

	"github.com/arielril/hktb/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Printf("failed to execute command: %s\n", err)
		os.Exit(1)
	}
}
