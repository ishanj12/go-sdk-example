# go-sdk-example

This example application starts a hello world HTTP server on port 8085 and then uses the [ngrok Go SDK](https://github.com/ngrok/ngrok-go) (`golang.ngrok.com/ngrok/v2`) to forward public traffic to that server. See the [quickstart](https://ngrok.com/docs/getting-started/go/) and [SDK reference](https://pkg.go.dev/golang.ngrok.com/ngrok/v2) for more details. When you run it, you'll get a public URL that anyone can use to access your app.

## Clone and Run This Example

```sh
git clone git@github.com:ngrok/go-sdk-example.git
cd go-sdk-example
NGROK_AUTHTOKEN=<token> go run main.go
```

## License

MIT
