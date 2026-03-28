// package main

// import (
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/maheshlode/learning_go/handlers"
// )

// func main() {

// 	// http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
// 	// 	log.Println("Good Bye")
// 	// 	d, err := io.ReadAll(r.Body)
// 	// 	if err != nil {
// 	// 		http.Error(rw, "Oops", http.StatusBadRequest)
// 	// 		return
// 	// 	}
// 	// 	fmt.Fprintf(rw, "Hello %s", d)
// 	// })

// 	// http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
// 	// 	log.Println("Hello World")
// 	// })

// 	l := log.New(os.Stdout, "product-api", log.LstdFlags)

// 	hh := handlers.NewHello()

// 	http.ListenAndServe("127.0.0.1:8080", nil)
// }

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/maheshlode/learning_go/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

	// http.ListenAndServe("127.0.0.1:8080", sm)
}
