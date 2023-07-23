package models

type Order struct {
	Id            int64   `json:"order_id"`
	CourierId     int64   
	Weight        float32 `json:"weight" binding:"required"`
	Cost          int32   `json:"cost" binding:"required"`
	Region        int32   `json:"regions" binding:"required"`
	DeliveryHours string  `json:"delivery_hours" binding:"required"`
	CompleteTime  string  `json:"complete_time,omitempty"`
}
