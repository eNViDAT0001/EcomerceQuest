package io

type OrderReport struct {
	Orders     uint
	Waiting    uint
	Delivering uint
	Delivered  uint
	Canceled   uint
	Revenue    uint
}
type OrderReportFilter struct {
	ProviderID uint
}
