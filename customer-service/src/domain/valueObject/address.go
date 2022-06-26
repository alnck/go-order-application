package valueobject

type Address struct {
	AddressLine string `json:"addressLine" bson:"addressLine"`
	City        string `json:"city" bson:"city" validate:"required"`
	Country     string `json:"country" bson:"country" validate:"required"`
	CityCode    int    `json:"cityCode" bson:"cityCode" validate:"required"`
}
