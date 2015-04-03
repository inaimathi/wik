package main

import (
	"fmt"
	"flag"
	"strings"
	"net/http"
)

func main() {
	var port = flag.Int("port", 8080, "Specify the TCP port this server should listen on. Defaults to 8080.")
	flag.Parse()

	wik := &Wiki{strings.TrimRight(flag.Arg(0), "/")}
	WikiHandlers(wik)

	fmt.Printf("Serving %s\n", flag.Arg(0))
	fmt.Printf("Listening on local port %d...\n", *port)
	
	static := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", static))

	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
