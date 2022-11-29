package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Please specify a filename")
		os.Exit(1)
	}

	filename := os.Args[1]
	if filename == "" {
		fmt.Fprintln(os.Stderr, "Please specify a filename")
		os.Exit(1)
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	if _, err := io.Copy(os.Stdout, file); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
