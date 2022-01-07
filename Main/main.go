package main

import (
	"fmt"
	"net/http"
)

func homePage(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(writer, "Home Page is here...")
	fmt.Println("Endpoint hit to Home Page...")
}

func aboutMe(writer http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(writer, "Pankaj Katyare is the best person in the world...")
	fmt.Println("Endpoint hit to Pankaj Katyare...")
}
func main() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutMe", aboutMe)

	fmt.Println("Server listen at port :8002...")
	http.ListenAndServe(":8002", nil)
}
