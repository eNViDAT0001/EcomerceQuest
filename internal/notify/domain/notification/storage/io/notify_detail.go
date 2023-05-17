package io

type NotificationDetail struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Seen    *bool  `json:"seen"`
	URL     string `json:"url"`
}
type NotificationInput struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Seen    *bool  `json:"seen"`
	URL     string `json:"url"`
	Image   string `json:"image"`
}
