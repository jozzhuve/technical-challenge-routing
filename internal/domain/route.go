package domain

// Route representa la mejor ruta encontrada desde una base hasta el lugar del accidente.
type Route struct {
	FromDepot string   `json:"fromDepot"`
	To        string   `json:"to"`
	Path      []string `json:"path"`
	Distance  float64  `json:"distance"`
}
