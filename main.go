package main

import (
	"fmt"
	"net/http"
	"text/template"
)

//config port
const (
	webServerPort = "9000"
	baseURL       = "0.0.0.0:" + webServerPort
)

func main() {
	http.HandleFunc("/", handleRoute)
	http.HandleFunc("/login", handleLogin)

	http.ListenAndServe(baseURL, nil)
}

func handleRoute(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("login_template.html")
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	//authenticate
	ok, data, err := AuthUsingLDAP(username, password)
	if !ok {
		http.Error(w, "invalid username/password", http.StatusUnauthorized)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	message := fmt.Sprintf("Welcome %s\n", data.FullName)
	w.Write([]byte(message))
}
