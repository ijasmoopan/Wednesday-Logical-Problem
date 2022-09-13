package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Request struct {
	Message string `json:"reqmessage"`
}
type Response struct {
	Message []string `json:"resmessage"`
}
var (
	word string
	str string
)

func main() {
	router := chi.NewRouter()
	router.Get("/hello", HelloHandler)
	router.Post("/post", PostHandler)
	http.ListenAndServe(":8000", router)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var request Request
	var response Response
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal("No request body", err)
	}
	log.Println("Request: ", request)
	flag1 := false
	flag2 := false
	fmt.Println("Message: ", request.Message)
	for idx, char := range request.Message {
		flag1 = true
		flag2 = true
		if char != ' ' {
			word = fmt.Sprint(word, string(char))
		} else {
			str = fmt.Sprint(str, word)
			flag2 = false
			word = string(char)
		}
		if idx % 10 == 0 && idx != 0 {
			response.Message = append(response.Message, str)
			flag1 = false
			str = ""
			fmt.Println("Response: ", response.Message)
		}
	}
	if flag1 {
		if flag2 {
			str = fmt.Sprint(str, word)
		}
		fmt.Println("String: ", str)
		response.Message = append(response.Message, str)
		fmt.Println("Response: ", response.Message)
	}
	json.NewEncoder(w).Encode(&response)
}

// hr@wednesday.is