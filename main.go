package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URLDescription struct {
	URL        string `json:"url"`
	Method     string `json:"method"`
	Desciption string `json:"description"`
	Payload    string `json:"payload,omitempty"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:        "/",
			Method:     "GET",
			Desciption: "See Documentation",
		},
		{
			URL:        "/blocks",
			Method:     "GET",
			Desciption: "Add a Block",
			Payload:    "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}
func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
