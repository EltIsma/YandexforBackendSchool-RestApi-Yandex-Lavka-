package dto


type OrderCompleteDto struct {
	CourierId         int    `json:"courier_id"`
	OrderId           int    `json:"order_id"`
	OrderCompleteTime string `json:"complete_time"`
}