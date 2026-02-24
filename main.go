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
	startServer()
	connectNgrok()
}

// This HTTP server is just for demonstration. If you already have an app
// running, skip startServer() and point ngrok.Forward() at its port instead.
func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		fmt.Fprintln(w, "Hello from ngrok-go!")
	})

	ln, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatal(err)
	}
	go http.Serve(ln, nil)
	log.Println("Server listening on port 8085")
}

func connectNgrok() {
	fwd, err := ngrok.Forward(context.Background(),
		ngrok.WithUpstream("http://localhost:8085"),

		// Uncomment below to use a specific domain.
		// https://dashboard.ngrok.com/domains
		// ngrok.WithURL("https://<your_domain_here>"),

		// Uncomment below to load balance across multiple instances of your app.
		// https://ngrok.com/docs/universal-gateway/endpoint-pooling/
		// ngrok.WithPoolingEnabled(true),

		// Uncomment below to require visitors to log in with Google before accessing your app.
		// https://ngrok.com/docs/traffic-policy/actions/oauth/
		// ngrok.WithTrafficPolicy(`
		//   on_http_request:
		//     - actions:
		//         - type: oauth
		//           config:
		//             provider: google
		// `),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Available at:", fwd.URL())

	select {}
}
