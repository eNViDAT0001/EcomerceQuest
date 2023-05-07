package io

type BannerUpdateReq struct {
	Title         string `json:"title"`
	Collection    string `json:"collection"`
	Discount      *int   `json:"discount"`
	Image         string `json:"image"`
	EndTime       string `json:"end_time"`
	ProductIDsIN  []uint `json:"product_ids_in"`
	ProductIDsOUT []uint `json:"product_ids_out"`
}
