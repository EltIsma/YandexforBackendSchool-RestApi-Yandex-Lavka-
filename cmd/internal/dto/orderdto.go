package dto

type OrderDto struct {
	Id            int    `json:"id"`
	Weight        float32 `json:"weight" binding:"required"`
	Cost          int32   `json:"cost" binding:"required"`
	Region        int32   `json:"regions" binding:"required"`
	DeliveryHours string  `json:"delivery_hours" binding:"required"`
	CompleteTime  string  `json:"complete_time,omitempty"`
}