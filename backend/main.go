package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	registerEndpoint := func(path string) {
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Max-Age", "86400")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			if r.Method != http.MethodPost {
				w.Header().Set("Allow", http.MethodPost)
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
				return
			}

			// Drain the body to allow clients to reuse connections.
			n, _ := io.Copy(io.Discard, r.Body)
			_ = r.Body.Close()

			log.Printf("request: method=%s path=%s size_bytes=%d remote=%s", r.Method, r.URL.Path, n, r.RemoteAddr)

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("{}"))
		})
	}

	registerEndpoint("/v1/traces")
	registerEndpoint("/v1/logs")
	registerEndpoint("/v1/metrics")

	server := &http.Server{
		Addr:    ":14318",
		Handler: mux,
	}

	log.Printf("listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
