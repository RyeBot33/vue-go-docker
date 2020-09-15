package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	FirstName string `json: "firstname"`
	LastName  string `json: "lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/write", write)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func index(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	p := Person{
		FirstName: "Ryan",
		LastName:  "Brothers",
		Email:     "test@test.com",
		Password:  "JennaNen",
	}
	bs, err := json.Marshal(p)
	if err != nil {
		error.Error(err)
	}
	fmt.Printf("%+v\n",p)
	fmt.Fprintf(w, string(bs))
}

func write(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	var data Person
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
//	bs, _ := json.Marshal(data) //To Print JSON
	fmt.Println(data.FirstName)
}
