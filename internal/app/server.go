package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/matkv/core/internal/config"
	"github.com/matkv/core/internal/random"
	"github.com/matkv/core/internal/ui"
)

func Start(port int) error {
	mux := http.NewServeMux()

	// 1. API Routes
	mux.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Go Core!"))
	})

	mux.HandleFunc("/api/random", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		value := random.Int(1000)
		_ = json.NewEncoder(w).Encode(map[string]int{"value": value})
	})

	// Settings: expose current configuration
	mux.HandleFunc("/api/settings", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(config.C)
	})

	// 2. Serve SvelteKit Frontend
	distFS, err := ui.GetDistFS()
	if err != nil {
		return fmt.Errorf("failed to load UI: %w", err)
	}

	// Handle static files
	fileServer := http.FileServer(http.FS(distFS))
	mux.Handle("/", fileServer)

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server starting on http://localhost%s\n", addr)
	return http.ListenAndServe(addr, mux)
}
