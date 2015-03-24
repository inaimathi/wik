package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Println(os.Args) // command-line args
	fmt.Println(os.Args[1:]) // as above, without program name
	
	// wik := &Wiki{"/home/inaimathi/tmp"}
	// wik.edit("test.md", []byte("Testing testing, one two three"))
	// wik.remove("test.md")
}
