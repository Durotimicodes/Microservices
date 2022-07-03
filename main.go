package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Durotimicodes/working/handlers"
)

func main() {

	//define a logger
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	//create an instrance of handlers
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)

	//create a new serve mux
	sm := http.NewServeMux()
	//register handlers
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)


	//manual server
	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1 *time.Second,
		WriteTimeout: 1 *time.Second,
	}

	//define the server
	log.Println("Server is running on port 9090...")
	s.ListenAndServe()


}
