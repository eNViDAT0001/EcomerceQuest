package io

type AddressPreview struct {
	ID       uint   `json:"id"`
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Gender   bool   `json:"gender"`
	Phone    string `json:"phone"`
	Province string `json:"province"`
	District string `json:"district"`
	Ward     string `json:"ward"`
	Street   string `json:"street"`
}
