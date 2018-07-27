package main

import "fmt"
import "net/http"
import "log"
import "encoding/json"

func main() {
	port:= 8080

	http.HandleFunc("/helloworld", helloWorldHandler)
	log.Printf("Server stating on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type helloWorldResponse struct {
// change the output field to be "messge"
         Message string `json:"message"`
// do not output this field
	Author string `json:"-"`
// do not output this field if the value is empty
	Date string `json:",omitempty"`
// convert output to a string and rename "id"
	Id int `json:"id, string"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse {Message: "Hello World"}
	data, err :=  json.Marshal(response)
	if err != nil {
	  panic("Ooops")
	}
	fmt.Fprint(w, string(data))
}


