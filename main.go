package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	// http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
	// 	log.Println("Good Bye")
	// 	d, err := io.ReadAll(r.Body)
	// 	if err != nil {
	// 		http.Error(rw, "Oops", http.StatusBadRequest)
	// 		return
	// 	}
	// 	fmt.Fprintf(rw, "Hello %s", d)
	// })

	// http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
	// 	log.Println("Hello World")
	// })

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello()

	http.ListenAndServe("127.0.0.1:8080", nil)
}
