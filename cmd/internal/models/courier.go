package models

import "time"

type Courier struct {
	Id           int64
	Type         string
	Regions      []int32  `pg:",array"`
	WorkingHours []string `pg:",array"`
}


const price = 1000

func (c *Courier) Earnings(count_completed_orders int) int {
	var earnings int
	if c.Type == "FOOT" {
		earnings = 2 * (price * count_completed_orders)
	} else if c.Type == "BIKE" {
		earnings = 3 * (price * count_completed_orders)
	} else {
		earnings = 4 * (price * count_completed_orders)
	}

	return earnings
}

func (c *Courier) Rating(count_completed_orders int, start_date string, end_date string) int {
	var coeff int
	if c.Type == "FOOT" {
		coeff = 2
	} else if c.Type == "BIKE" {
		coeff = 3
	} else {
		coeff = 4
	}
	startT, err := time.Parse("2006-01-02", start_date)
	if err != nil {
		return 0
	}
	endT, err := time.Parse("2006-01-02", end_date)
	if err != nil {
		return 0
	}
	duration := endT.Sub(startT)
	hours := int(duration.Hours())
	rating := (count_completed_orders / hours) * coeff
	return rating
}
