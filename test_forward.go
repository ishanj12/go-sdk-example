//go:build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.ngrok.com/ngrok/v2"
)

func main() {
	// Simple server on port 8080 (simulates "existing app")
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello from existing Go app!")
		})
		log.Println("Existing app on http://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// --- Snippet from README "Add to Existing Code" ---
	if err := forwardToApp(); err != nil {
		log.Fatal(err)
	}
	select {} // keep alive
}

func forwardToApp() error {
	fwd, err := ngrok.Forward(context.Background(),
		ngrok.WithUpstream("localhost:8080"),
	)
	if err != nil {
		return err
	}
	log.Println("Ingress established at:", fwd.URL())
	return nil
}
