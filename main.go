package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.ngrok.com/ngrok/v2"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	// Start a simple HTTP server
	go http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		fmt.Fprintln(w, "Hello from go-sdk-example!")
	}))

	// Forward ngrok traffic to the local server
	fwd, err := ngrok.Forward(ctx,
		ngrok.WithUpstream("localhost:8080"),
	)
	if err != nil {
		return err
	}

	log.Println("Ingress established at:", fwd.URL())
	<-fwd.Done()
	return nil
}
