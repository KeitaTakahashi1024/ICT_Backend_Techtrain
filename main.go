package main

import(
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/user/create", createUser)
	http.HandleFunc("/user/get", getUser)
	http.HandleFunc("/user/update", updateUser)
	http.ListenAndServe(":8080", nil)
}




func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to user/create, from Docker container!")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to user/get, from Docker container!")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to user/update, from Docker container!")
}
