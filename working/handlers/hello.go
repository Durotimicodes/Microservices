package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello{
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	dataSlice, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ooops something went wrong!", http.StatusBadRequest)
		return
	}
	h.l.Printf("Data %s\n", dataSlice)

	fmt.Fprintf(w, "Hello %s", dataSlice)
}
