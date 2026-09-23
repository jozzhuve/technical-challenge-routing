package domain

import "fmt"

// Graph representa una red dirigida de nodos donde cada arista tiene una distancia no negativa.
type Graph map[string]map[string]float64

// Validate verifica que el grafo tenga contenido y que todas las distancias sean válidas.
func (g Graph) Validate() error {
	if len(g) == 0 {
		return fmt.Errorf("el grafo no puede estar vacío")
	}

	for from, edges := range g {
		if from == "" {
			return fmt.Errorf("el nombre de un nodo no puede estar vacío")
		}
		for to, distance := range edges {
			if to == "" {
				return fmt.Errorf("el nombre de un nodo destino no puede estar vacío")
			}
			if distance < 0 {
				return fmt.Errorf("la distancia entre %s y %s no puede ser negativa", from, to)
			}
		}
	}

	return nil
}
