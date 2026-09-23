package httpapi

import "net/http"

// NewServer construye el router HTTP del servicio y registra endpoints versionados y de salud.
func NewServer(handler *RouteHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/routes/optimal", handler.Calculate)
	mux.HandleFunc("POST /routes/optimal", handler.Calculate)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "UP"})
	})
	return mux
}
