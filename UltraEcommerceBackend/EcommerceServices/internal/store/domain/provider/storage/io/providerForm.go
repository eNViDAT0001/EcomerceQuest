package io

type ProviderForm struct {
	ID        uint
	UserID    uint
	Name      string
	ImagePath string
}
type ProviderUpdateForm struct {
	Name      string
	ImagePath string
}
