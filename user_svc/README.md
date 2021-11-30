### Code structure
todo/
|---cmd/
|------service/
|----------server.go          Wire the service.
|----------server_gen.go      Also wire the service.
|------main.go                Runs the service
|---pkg/
|------endpoints/
|----------endpoint.go        The endpoint logic.
|----------endpoint_gen.go    This will wire the endpoints.
|----------middleware.go      Endpoint middleware
|------http/
|----------handler.go         Transport logic encode/decode data.
|----------handler_gen.go     This will wire the transport.
|------io/
|----------io.go              The input output structs.
|------service/
|----------middleware.go      The service middleware.
|----------service.go         Business logic.
