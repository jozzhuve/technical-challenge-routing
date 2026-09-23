package application_test

import (
	"errors"
	"testing"

	"github.com/jozzhuve/technical-challenge-routing/internal/algorithm"
	"github.com/jozzhuve/technical-challenge-routing/internal/application"
	"github.com/jozzhuve/technical-challenge-routing/internal/domain"
)

// TestExecuteChoosesNearestDepot valida que se seleccione la base con menor distancia total.
func TestExecuteChoosesNearestDepot(t *testing.T) {
	service := application.NewCalculateOptimalRouteService(algorithm.NewDijkstra())
	graph := domain.Graph{
		"Miraflores": {"San Isidro": 7},
		"Ate":        {"Surco": 10},
		"Surco":      {"San Isidro": 6},
		"San Isidro": {},
	}

	result, err := service.Execute(application.CalculateRouteCommand{
		AccidentLocation: "San Isidro",
		Depots:           []string{"Ate", "Miraflores"},
		Graph:            graph,
	})

	if err != nil {
		t.Fatalf("no se esperaba error: %v", err)
	}
	if result.FromDepot != "Miraflores" || result.Distance != 7 {
		t.Fatalf("se esperaba Miraflores a distancia 7 y se obtuvo %+v", result)
	}
}

// TestExecuteReturnsControlledError valida el escenario donde ninguna base llega al accidente.
func TestExecuteReturnsControlledError(t *testing.T) {
	service := application.NewCalculateOptimalRouteService(algorithm.NewDijkstra())
	graph := domain.Graph{
		"Ate":        {},
		"San Isidro": {},
	}

	_, err := service.Execute(application.CalculateRouteCommand{
		AccidentLocation: "San Isidro",
		Depots:           []string{"Ate"},
		Graph:            graph,
	})

	if !errors.Is(err, application.ErrUnreachable) {
		t.Fatalf("se esperaba ErrUnreachable y se obtuvo %v", err)
	}
}
