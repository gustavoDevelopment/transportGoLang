package catalog

type TruckBrand struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type TruckBrands []TruckBrand
