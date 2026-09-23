package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jozzhuve/technical-challenge-routing/internal/application"
	"github.com/jozzhuve/technical-challenge-routing/internal/domain"
)

// RouteRequest representa el contrato HTTP recibido para calcular la ruta óptima.
type RouteRequest struct {
	AccidentLocation string       `json:"accidentLocation"`
	Depots           []string     `json:"depots"`
	Graph            domain.Graph `json:"graph"`
}

// RouteHandler adapta solicitudes HTTP al caso de uso de cálculo de rutas.
type RouteHandler struct {
	service *application.CalculateOptimalRouteService
}

// NewRouteHandler crea el handler con el caso de uso requerido.
func NewRouteHandler(service *application.CalculateOptimalRouteService) *RouteHandler {
	return &RouteHandler{service: service}
}

// Calculate procesa la entrada JSON y devuelve la ruta de menor distancia.
func (h *RouteHandler) Calculate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var request RouteRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&request); err != nil {
		writeError(w, http.StatusBadRequest, "INVALID_REQUEST", "El JSON enviado no es válido.")
		return
	}

	result, err := h.service.Execute(application.CalculateRouteCommand{
		AccidentLocation: request.AccidentLocation,
		Depots:           request.Depots,
		Graph:            request.Graph,
	})
	if err != nil {
		if errors.Is(err, application.ErrUnreachable) {
			writeError(w, http.StatusUnprocessableEntity, "ROUTE_NOT_FOUND", err.Error())
			return
		}
		writeError(w, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}

	writeJSON(w, http.StatusOK, result)
}

// writeJSON serializa una respuesta JSON con el código HTTP indicado.
func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

// writeError construye el contrato homogéneo utilizado para errores controlados.
func writeError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, map[string]string{
		"code":    code,
		"message": message,
	})
}
