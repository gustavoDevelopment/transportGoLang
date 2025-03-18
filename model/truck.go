package model

type Truck struct {
	LicensePlate  string `json:"license_plate"`  // Placa del camión
	Brand         string `json:"brand"`          // Marca del camión (Ej: Volvo, Scania)
	Model         string `json:"model"`          // Modelo del camión
	Year          int    `json:"year"`           // Año de fabricación
	Color         string `json:"color"`          // Color del camión
	OwnerName     string `json:"owner_name"`     // Nombre del propietario
	Capacity      int    `json:"capacity_kg"`    // Capacidad de carga en kg
	ChassisNumber string `json:"chassis_number"` // Número de chasis
	EngineNumber  string `json:"engine_number"`  // Número de motor
}
