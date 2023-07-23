package dto


type CourierDtoType string

const (
	Foot CourierDtoType = "FOOT"
	Bike CourierDtoType = "BIKE"
	Auto CourierDtoType = "AUTO"
)

type CourierDto struct {
	Id           int64    `json:"courier_id"`
	Type         CourierDtoType   `json:"courier_type" binding:"required"`
	Regions      []int32  `json:"regions" binding:"required"`
	WorkingHours []string `json:"working_hours" binding:"required"`
}
type CourierDtoEarningRating struct {
	Earning       int      `json:"earnings,omitempty"`
	Ratings       int      `json:"rating,omitempty"`
}