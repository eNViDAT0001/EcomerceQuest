package io

type SeenMessageEvent struct {
	ID     uint `json:"id"`
	UserID uint `json:"user_id"`
	ToID   uint `json:"to_id"`
}
