package main

import (
	"fmt"
	"net/http"
	"os"
)

func quitquitquit(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	os.Exit(0)
}

func health(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/quitquitquit", quitquitquit)
	http.HandleFunc("/healthz/ready", health)
	err := http.ListenAndServe(":15020", nil)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}
