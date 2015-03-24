package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Println(os.Args) // command-line args
	fmt.Println(os.Args[1:]) // as above, without program name
	
	wik := &Wiki{"/home/inaimathi/tmp"}
	wik.Create("test.md")
	wik.Edit("test.md", []byte("### Testing testing, one two three"))
	body, _ := wik.Render("test.md")
	fmt.Println(string(body))
	body, _ = wik.Raw("test.md")
	fmt.Println(string(body))
	wik.Remove("test.md")
}
