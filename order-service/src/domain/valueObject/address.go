package valueobject

type Address struct {
	AddressLine string `json:"addressLine" bson:"addressLine"`
	City        string `json:"city" bson:"city"`
	Country     string `json:"country" bson:"country"`
	CityCode    int    `json:"cityCode" bson:"cityCode"`
}
