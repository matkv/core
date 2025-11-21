package server

import (
	"fmt"
	"net/http"

	"github.com/matkv/core/internal/ui"
)

func Start(port int) error {
	mux := http.NewServeMux()

	// 1. API Routes
	mux.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Go Core!"))
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
