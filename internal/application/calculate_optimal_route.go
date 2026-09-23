package application

import (
	"errors"
	"fmt"
	"math"

	"github.com/jozzhuve/technical-challenge-routing/internal/algorithm"
	"github.com/jozzhuve/technical-challenge-routing/internal/domain"
)

// ErrUnreachable indica que ninguna base puede alcanzar la ubicación del accidente.
var ErrUnreachable = errors.New("la ubicación del accidente no es alcanzable desde ninguna base")

// CalculateRouteCommand contiene los datos necesarios para buscar la mejor base y ruta.
type CalculateRouteCommand struct {
	AccidentLocation string
	Depots           []string
	Graph            domain.Graph
}

// CalculateOptimalRouteService coordina la búsqueda del camino mínimo para múltiples bases.
type CalculateOptimalRouteService struct {
	dijkstra *algorithm.Dijkstra
}

// NewCalculateOptimalRouteService crea el caso de uso con el algoritmo de Dijkstra.
func NewCalculateOptimalRouteService(dijkstra *algorithm.Dijkstra) *CalculateOptimalRouteService {
	return &CalculateOptimalRouteService{dijkstra: dijkstra}
}

// Execute valida la entrada, evalúa cada base y devuelve la ruta global de menor distancia.
func (s *CalculateOptimalRouteService) Execute(command CalculateRouteCommand) (domain.Route, error) {
	if err := validateCommand(command); err != nil {
		return domain.Route{}, err
	}

	bestDistance := math.Inf(1)
	var bestRoute domain.Route
	found := false

	for _, depot := range command.Depots {
		path, distance, reachable := s.dijkstra.ShortestPath(command.Graph, depot, command.AccidentLocation)
		if !reachable {
			continue
		}
		if !found || distance < bestDistance {
			bestDistance = distance
			found = true
			bestRoute = domain.Route{
				FromDepot: depot,
				To: command.AccidentLocation,
				Path: path,
				Distance: distance,
			}
		}
	}

	if !found {
		return domain.Route{}, ErrUnreachable
	}

	return bestRoute, nil
}

// validateCommand valida las precondiciones necesarias antes de ejecutar el algoritmo.
func validateCommand(command CalculateRouteCommand) error {
	if command.AccidentLocation == "" {
		return fmt.Errorf("accidentLocation es requerido")
	}
	if len(command.Depots) == 0 {
		return fmt.Errorf("se requiere al menos una base de grúas")
	}
	if err := command.Graph.Validate(); err != nil {
		return err
	}
	return nil
}
