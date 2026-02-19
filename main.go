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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		fmt.Fprintln(w, "Hello from ngrok-go!")
	})

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	go http.Serve(ln, nil)

	fwd, err := ngrok.Forward(context.Background(),
		ngrok.WithUpstream("http://localhost:8080"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Ingress established at:", fwd.URL())

	select {}
}
