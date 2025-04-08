package proxy

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// ProxyHandler handles incoming requests and forwards them to the appropriate backend services.
type ProxyHandler struct {
	// Add any necessary fields, such as a backend URL or a logger
}

// HandleRequest is a placeholder function to handle proxy requests.
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request handled successfully"))
}

// NewHandler creates and returns a new HTTP handler for the proxy.
func NewHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Proxy handler response"))
	})
}

// NewProxyHandler creates a new instance of ProxyHandler.
func NewProxyHandler() *ProxyHandler {
	return &ProxyHandler{}
}

// ServeHTTP processes incoming HTTP requests.
func (h *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Implement request processing and forwarding logic here
	log.Printf("Received request: %s %s", r.Method, r.URL)

	// Example: Forward the request to a backend service
	// backendURL := "http://backend-service" + r.URL.Path
	// resp, err := http.Get(backendURL)
	// if err != nil {
	//     http.Error(w, "Failed to reach backend service", http.StatusBadGateway)
	//     return
	// }
	// defer resp.Body.Close()
	// w.WriteHeader(resp.StatusCode)
	// io.Copy(w, resp.Body)

	// Placeholder response for demonstration
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request processed"))
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Helios Proxy"))
	})
	// Set a custom 404 handler
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})
	return router
}
