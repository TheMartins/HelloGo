package main

import "fmt"
import "net/http"
import "log"

func main() {
	port:= 8080

	http.HandleFunc("/helloworld", helloWorldHandler)
	log.Printf("Server stating on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}
