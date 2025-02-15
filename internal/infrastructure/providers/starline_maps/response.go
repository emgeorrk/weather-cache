package starline_maps

type GeocoderResponse struct {
	ID          int        `json:"osm_id"`
	Type        string     `json:"osm_type"`
	Lat         float64    `json:"lat"`
	Lon         float64    `json:"lon"`
	DisplayName string     `json:"display_name"`
	Address     OSMAddress `json:"address"`
	Distance    float64    `json:"distance"`
}

type OSMAddress struct {
	Name        string `json:"name"`
	HouseNumber string `json:"house_number"`
	Street      string `json:"street"`
	City        string `json:"city"`
	CityType    string `json:"city_type"`
	District    string `json:"district"`
	Region      string `json:"region"`
	Postcode    string `json:"postcode"`
	Country     string `json:"country"`
}
