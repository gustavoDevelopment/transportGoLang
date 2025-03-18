package catalog

type City struct {
	ID           string      `json:"id"`
	Code         string      `json:"code"`                 // Código de la ciudad
	Name         string      `json:"name"`                 // Nombre de la ciudad
	DepartmentID string      `json:"department_id"`        // Relación con el departamento
	Department   *Department `json:"department,omitempty"` // Relación con el departamento (objeto anidado)
}

type Citys []City
