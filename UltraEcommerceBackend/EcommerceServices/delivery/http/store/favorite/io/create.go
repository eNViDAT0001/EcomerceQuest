package io

type CreateFavoriteReq struct {
	UserID     uint `json:"user_id"`
	ProviderID uint `json:"provider_id"`
}
