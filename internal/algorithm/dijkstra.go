package algorithm

import (
	"container/heap"
	"math"

	"github.com/jozzhuve/technical-challenge-routing/internal/domain"
)

type queueItem struct {
	node     string
	distance float64
	index    int
}

type priorityQueue []*queueItem

func (q priorityQueue) Len() int { return len(q) }
func (q priorityQueue) Less(i, j int) bool { return q[i].distance < q[j].distance }
func (q priorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}
func (q *priorityQueue) Push(value any) {
	item := value.(*queueItem)
	item.index = len(*q)
	*q = append(*q, item)
}
func (q *priorityQueue) Pop() any {
	old := *q
	last := len(old) - 1
	item := old[last]
	old[last] = nil
	*q = old[:last]
	return item
}

// Dijkstra calcula el camino mínimo entre dos nodos usando pesos no negativos.
type Dijkstra struct{}

// NewDijkstra crea una instancia del algoritmo de camino mínimo.
func NewDijkstra() *Dijkstra {
	return &Dijkstra{}
}

// ShortestPath calcula la distancia y el recorrido mínimo entre origen y destino.
// El booleano retornado indica si el destino es alcanzable desde el origen.
func (d *Dijkstra) ShortestPath(graph domain.Graph, source, target string) ([]string, float64, bool) {
	if source == target {
		return []string{source}, 0, true
	}

	distances := make(map[string]float64, len(graph))
	previous := make(map[string]string, len(graph))
	for node := range graph {
		distances[node] = math.Inf(1)
	}
	distances[source] = 0

	queue := &priorityQueue{}
	heap.Init(queue)
	heap.Push(queue, &queueItem{node: source, distance: 0})

	for queue.Len() > 0 {
		current := heap.Pop(queue).(*queueItem)
		knownDistance, exists := distances[current.node]
		if exists && current.distance > knownDistance {
			continue
		}
		if current.node == target {
			break
		}

		for neighbor, weight := range graph[current.node] {
			candidate := current.distance + weight
			known, exists := distances[neighbor]
			if !exists || candidate < known {
				distances[neighbor] = candidate
				previous[neighbor] = current.node
				heap.Push(queue, &queueItem{node: neighbor, distance: candidate})
			}
		}
	}

	distance, reachable := distances[target]
	if !reachable || math.IsInf(distance, 1) {
		return nil, 0, false
	}

	path := buildPath(previous, source, target)
	if len(path) == 0 {
		return nil, 0, false
	}

	return path, distance, true
}

// buildPath reconstruye el camino desde el destino siguiendo los nodos predecesores.
func buildPath(previous map[string]string, source, target string) []string {
	path := []string{target}
	current := target

	for current != source {
		parent, exists := previous[current]
		if !exists {
			return nil
		}
		current = parent
		path = append(path, current)
	}

	for left, right := 0, len(path)-1; left < right; left, right = left+1, right-1 {
		path[left], path[right] = path[right], path[left]
	}
	return path
}
