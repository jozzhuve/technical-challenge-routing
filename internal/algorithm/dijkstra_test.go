package algorithm_test

import (
	"reflect"
	"testing"

	"github.com/jozzhuve/technical-challenge-routing/internal/algorithm"
	"github.com/jozzhuve/technical-challenge-routing/internal/domain"
)

// TestShortestPath valida el ejemplo principal del reto técnico.
func TestShortestPath(t *testing.T) {
	graph := domain.Graph{
		"Miraflores": {"San Isidro": 7, "Barranco": 3},
		"San Isidro": {"Miraflores": 7, "Lince": 4},
		"Barranco":   {"Miraflores": 3, "Surco": 5},
		"Lince":      {"San Isidro": 4, "Surco": 6},
		"Surco":      {"Barranco": 5, "Lince": 6, "Ate": 10},
		"Ate":        {"Surco": 10},
	}

	path, distance, reachable := algorithm.NewDijkstra().ShortestPath(graph, "Miraflores", "San Isidro")

	if !reachable {
		t.Fatal("se esperaba una ruta alcanzable")
	}
	if distance != 7 {
		t.Fatalf("se esperaba distancia 7 y se obtuvo %v", distance)
	}
	expected := []string{"Miraflores", "San Isidro"}
	if !reflect.DeepEqual(path, expected) {
		t.Fatalf("se esperaba %v y se obtuvo %v", expected, path)
	}
}
