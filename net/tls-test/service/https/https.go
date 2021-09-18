package main

import (
	"fmt"
	"net/http"
)

func handlerHttps(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of https service in golang!   liuxin")
}

func main() {
	http.HandleFunc("/", handlerHttps)
	err := http.ListenAndServeTLS(":8080", "/Users/dong/Developer/workspace/golang/tls-test/shell/server.crt",
		"/Users/dong/Developer/workspace/golang/tls-test/shell/server.key", nil)

	fmt.Println(err)
}