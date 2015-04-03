package main

import (
	"os"
	"fmt"
	"strings"
	"net/http"
)

func main() {
	fmt.Println(os.Args) // command-line args
	fmt.Println(os.Args[1:]) // as above, without program name
	fmt.Println(os.Args[1])

	wik := &Wiki{strings.TrimRight(os.Args[1], "/")}

	WikiHandlers(wik)
	
	static := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", static))

	http.ListenAndServe(":8080", nil)
}
