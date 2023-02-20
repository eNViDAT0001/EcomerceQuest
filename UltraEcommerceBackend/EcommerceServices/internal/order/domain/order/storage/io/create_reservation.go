package io

import "time"

type CreateReservationForm struct {
	OptionID uint
	OrderID  uint
	Quantity int
	EndTime  time.Time
}
