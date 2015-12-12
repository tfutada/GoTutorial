package main

import (
	"net/http"
	"encoding/json"
)

type User struct {
	Name string
	Age int
}

func main() {

	http.HandleFunc("/getUser", getUser)

	http.ListenAndServe(":8080", nil)
}

func getUser(w http.ResponseWriter, r *http.Request) {

	user := User {"David", 28}

	enc := json.NewEncoder(w)
	enc.Encode(user)
}