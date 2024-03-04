package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	Id    string
	Name  string
	Age int
}

var data = []student{
	{"A001", "Alfredo", 17},
	{"Z002", "Aziz", 18},
	{"U003", "Fajar", 19},
	{"C004", "Rohman", 17},
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)

}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var id = r.FormValue("id")
		var result []byte
		var err error

		for _, each := range data {
			if each.Id == id {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}

		}

		http.Error(w, "User not found", http.StatusNotFound)
		return

	}
	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("starting web server at localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
