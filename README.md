# Helios Proxy

Helios Proxy is a lightweight proxy server designed to handle incoming requests and route them to appropriate backend services while providing robust authorization features. This project implements both JWT and OAuth 2.0 for secure authentication.

## Features

- **Proxy Handling**: Efficiently forwards requests to backend services.
- **Authorization**: Supports JWT and OAuth 2.0 for secure access control.
- **Middleware**: Includes logging, authentication, and authorization checks.
- **End-to-End Testing**: Comprehensive tests to ensure functionality and security.

## Project Structure

```
helios-proxy
├── cmd
│   └── main.go            # Entry point of the application
├── pkg
│   ├── proxy
│   │   ├── handler.go     # Proxy handler implementation
│   │   └── middleware.go   # Middleware functions
│   ├── auth
│   │   ├── jwt.go         # JWT functions for token management
│   │   └── oauth.go       # OAuth 2.0 authorization flows
│   └── utils
│       └── helpers.go     # Utility functions
├── tests
│   ├── e2e
│   │   ├── proxy_test.go   # End-to-end tests for proxy functionality
│   │   └── auth_test.go    # End-to-end tests for authentication features
│   └── unit
│       ├── proxy_handler_test.go # Unit tests for proxy handler
│       └── auth_jwt_test.go      # Unit tests for JWT functions
├── go.mod                  # Go module definition
├── go.sum                  # Dependency checksums
└── README.md               # Project documentation
```

## Setup Instructions

1. Clone the repository:
   ```
   git clone <repository-url>
   cd helios-proxy
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Run the application:
   ```
   go run cmd/main.go
   ```

## Usage

- The proxy server listens for incoming requests and routes them based on defined rules.
- Authorization is handled through JWT or OAuth 2.0, depending on the configuration.

## Testing

To run the tests, use the following command:

```
go test ./...
```

This will execute both unit and end-to-end tests to ensure the functionality of the proxy and authorization features.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.