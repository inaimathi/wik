package main

import (
	"os"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(os.Args) // command-line args
	fmt.Println(os.Args[1:]) // as above, without program name
	
	wik := &Wiki{"/home/inaimathi/tmp"}

	WikiHandlers(wik)
	http.ListenAndServe(":8080", nil)
}
