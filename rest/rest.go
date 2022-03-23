package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mangosteen0903/go-coin/blockchain"
	"github.com/mangosteen0903/go-coin/utils"
)

var port string

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://223.130.161.220%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL        url    `json:"url"`
	Method     string `json:"method"`
	Desciption string `json:"description"`
	Payload    string `json:"payload,omitempty"`
}

type addBlockBody struct {
	Message string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:        url("/"),
			Method:     "GET",
			Desciption: "See Documentation",
		},
		{
			URL:        url("/blocks"),
			Method:     "GET",
			Desciption: "Add a Block",
			Payload:    "data:string",
		},
		{
			URL:        url("/blocks/{height}"),
			Method:     "GET",
			Desciption: "See a Block",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addblockbody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addblockbody))
		blockchain.GetBlockchain().AddBlock(addblockbody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["height"])
	utils.HandleErr(err)
	block := blockchain.GetBlockchain().GetBlock(id)
	json.NewEncoder(rw).Encode(block)
	fmt.Println(id)
}
func Start(aPort int) {
	router := mux.NewRouter()
	port = fmt.Sprintf(":%d", aPort)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/block/{height:[0-9]+}", block).Methods("GET")
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
