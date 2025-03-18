package catalog

type Department struct {
	ID   string `json:"id"`
	Code string `json:"code"` // Código del departamento
	Name string `json:"name"` // Nombre del departamento
}

type Departments []Department
