package main

import (
	"net/http"
	"fmt"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	//fmt.Println(writer, "Hello World, %s", request.URL.Path[1:])
	fmt.Println(request)
	fmt.Println()
	fmt.Fprintf(writer, "Hello World, %s", request.URL.Path[1:])
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
