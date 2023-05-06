package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
func getGoogleFeedback(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://google.com/")
	if err != nil {
		// handle error
		io.WriteString(w, "Google No Found!\n")
	}
	defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
		io.WriteString(w, "Internal Error!\n")
	}
	fmt.Printf("got /google request\n")
	io.WriteString(w, string(body))
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/google", getGoogleFeedback)

	err := http.ListenAndServe(":3333", nil)
    if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
