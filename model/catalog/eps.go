package catalog

type EPS struct {
	ID   string `json:"id"`
	Code string `json:"code"` // Código de la EPS
	Name string `json:"name"` // Nombre de la EPS
}

type EPSs []EPS
