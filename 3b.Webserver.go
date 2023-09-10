package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port string

	fmt.Print("Enter the port number to listen on: ")
	fmt.Scan(&port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Friends!")
	})

	log.Println("Server listening on port", port, "....")
	http.ListenAndServe(":"+port, nil)
}

// Output
// Enter the port number to listen on: 5020
// 2023/09/10 05:59:32 Server listening on port 5020 ....
