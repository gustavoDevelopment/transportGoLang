package catalog

type DocumentType struct {
	ID   string `json:"id"`
	Code string `json:"code"` // CC, TI, NIT, PAS
	Name string `json:"name"` // Cédula de Ciudadanía, Tarjeta de Identidad, etc.
}

type DocumentTypes []DocumentType
