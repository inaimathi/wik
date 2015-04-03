package main

import (
	"os"
	"fmt"
	"strings"
	"net/http"
)

func main() {
	wik := &Wiki{strings.TrimRight(os.Args[1], "/")}
	WikiHandlers(wik)

	fmt.Print("Serving ")
	fmt.Println(os.Args[1])
	fmt.Println("Listening on local port 8080...")	
	
	static := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", static))

	http.ListenAndServe(":8080", nil)
}
