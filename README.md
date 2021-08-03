# Hexagonal Architecture - Sample App.

### Server

> App listens on port 7777 (default)

- Rest APIs
    - POST `http://{hostname:port}/beers` 
    - POST `http://{hostname:port}/reviews`
    

### Domain Area

> service to manage beers & to manage reviews for them

- Beer
- Review(for Beer)
    

### Components

- Ports
    - `pkg/core/beer/repository.go`
    - `pkg/core/review/repository.go`

- Core
    - Model
        - `pkg/core/beer/beer.go`
        - `pkg/core/review/review.go`
    - Biz
        - `pkg/core/beer/service.go`
        - `pkg/core/review/service.go`

- Adapters
    - repositories
        - `pkg/repository/*`
    - http
        - `pkg/http/*`

- Dependency Injection
    - `cmd/server/main.go`
