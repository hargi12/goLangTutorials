// shows how to read user data from a html form using gorilla schema package
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

// create a new struct to holder the user details
type User struct {
	Username string
	Password string
}

// create a function to read form data

func readForm(r *http.Request) *User {
	r.ParseForm()
	user := new(User)
	decoder := schema.NewDecoder()
	decodeErr := decoder.Decode(user, r.PostForm)
	if decodeErr != nil {
		log.Printf("Error mapping parsed form data to struct : ", decodeErr)
	}
	return user
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parsedTemplate, _ := template.ParseFiles("templates/login-form.html")
		parsedTemplate.Execute(w, nil)
	} else {
		user := readForm(r)
		fmt.Fprintf(w, "Hello "+user.Username+"!")
	}
}

func main() {
	http.HandleFunc("/", login)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("Error starting http server : ", err)
		return
	}

}
