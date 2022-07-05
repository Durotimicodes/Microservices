package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
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
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//define the server
	log.Println("Server is running on port 9090...")

	//wrapping up server in a go function to prevent it from blocking
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	sigChan := make(chan os.Signal)

	/*this will broadcast a message on this channel whenever an operating
	system kill command OR whenever an operating system interrupt
	*/
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	//sig receiving the sig channel
	sig := <- sigChan
	l.Println("Received terminate, gracefully shudown", sig)

	//i want you to allow 30 secs to gracefully shutdown
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	//graceful shutdown
	s.Shutdown(tc)

}
