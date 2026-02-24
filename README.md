# go-sdk-example

This example application starts a hello world HTTP server on port 8085 and then uses the [ngrok Go SDK](https://github.com/ngrok/ngrok-go) (`golang.ngrok.com/ngrok/v2`) to forward public traffic to that server. See the [quickstart](https://ngrok.com/docs/getting-started/go/) and [SDK reference](https://pkg.go.dev/golang.ngrok.com/ngrok/v2) for more details. When you run it, you'll get a public URL that anyone can use to access your app.

## Option 1: Clone and Run This Example

```sh
git clone git@github.com:ngrok/go-sdk-example.git
cd go-sdk-example
NGROK_AUTHTOKEN=<token> go run main.go
```

## Option 2: Add ngrok to Your Own App

1. Install the SDK:

   ```sh
   go get golang.ngrok.com/ngrok/v2
   ```

2. Add the following to your app:

   ```go
   import "golang.ngrok.com/ngrok/v2"

   func connectNgrok() {
       fwd, err := ngrok.Forward(context.Background(),
           ngrok.WithUpstream("http://localhost:8085"),
       )
       if err != nil {
           log.Fatal(err)
       }
       log.Println("Available at:", fwd.URL())
       select {}
   }
   ```

3. Set your authtoken:

   ```sh
   export NGROK_AUTHTOKEN=<token>
   ```

## License

MIT
