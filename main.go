package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.ngrok.com/ngrok/v2"
)

func main() {
	fwd, err := ngrok.Forward(context.Background(),
		ngrok.WithUpstream("localhost:8080"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Ingress established at:", fwd.URL())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		fmt.Fprintln(w, "Hello from ngrok-go!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
