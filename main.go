package main

import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	fmt.Println(os.Args) // command-line args
	fmt.Println(os.Args[1:]) // as above, without program name

	wik := &Wiki{"/home/inaimathi/tmp"}

	WikiHandlers(wik)
	
	static := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", static))

	http.ListenAndServe(":8080", nil)
}
