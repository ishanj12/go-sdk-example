//go:build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"golang.ngrok.com/ngrok/v2"
)

func main() {
	// Simple server on port 8080 (simulates "existing app")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from existing Go app!")
	})
	ln, _ := net.Listen("tcp", ":8080")
	go http.Serve(ln, nil)
	log.Println("Existing app on http://localhost:8080")

	// --- Snippet from README "Add to Existing Code" ---
	fwd, err := ngrok.Forward(context.Background(),
		ngrok.WithUpstream("http://localhost:8080"),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Ingress established at:", fwd.URL())

	select {} // keep alive — fwd stays in scope
}
