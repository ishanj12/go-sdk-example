# go-sdk-example

A minimal HTTP server using the [ngrok Go SDK](https://github.com/ngrok/ngrok-go) (`golang.ngrok.com/ngrok/v2`).

## Clone and Run

```sh
git clone git@github.com:ngrok/go-sdk-example.git
cd go-sdk-example
NGROK_AUTHTOKEN=<token> go run main.go
```

## Add to Existing Code

1. Install the SDK:

   ```sh
   go get golang.ngrok.com/ngrok/v2
   ```

2. Set your authtoken:

   ```sh
   export NGROK_AUTHTOKEN=<token>
   ```

3. Add the following to your app:

   ```go
   import "golang.ngrok.com/ngrok/v2"

   func forwardToApp() error {
       // Reads NGROK_AUTHTOKEN from environment automatically
       fwd, err := ngrok.Forward(context.Background(),
           ngrok.WithUpstream("http://localhost:8080"),
       )
       if err != nil {
           return err
       }
       log.Println("Ingress established at:", fwd.URL())
       return nil
   }
   ```

## License

MIT
