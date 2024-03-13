package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Response struct {
	Error error `json:"error"`
	Data  User  `json:"data"`
}

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, web")
	})
	http.HandleFunc("GET /greet/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		fmt.Fprintf(w, "Hello, %s!", name)
	})
	http.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		var err error
		res := Response{}

		buff, err := io.ReadAll(r.Body)

		var user User
		err = json.Unmarshal(buff, &user)
		if err != nil {
			res.Error = err
		}

		res.Data = user

		resData, err := json.Marshal(res)
		if err != nil {
			log.Println(err)
			return
		}

		w.Header().Add("type", "application/json")
		fmt.Fprint(w, string(resData))
	})

	http.ListenAndServe(":8080", nil)
}
