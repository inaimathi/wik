package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Println(os.Args) // command-line args
	fmt.Println(os.Args[1:]) // as above, without program name
	p1 := &Page{Path: "Welcome.md", Body: []byte("# This is a simple index page...")}
	p1.save()
	p2, _ := loadPage("Welcome.md")
	fmt.Println(string(p2.render()))
	file, _ := os.Open("Welcome.md")
	fstat, _ := file.Stat()
	fmt.Println(fstat.ModTime().String())
	file.Close()
}

// func main() {
// 	fmt.Println("Testing...")
// 	fmt.Println("And, testing again and again")
// }
