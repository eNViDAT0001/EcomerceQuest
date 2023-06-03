package io

type UpdateRequestContactForm struct {
	Subject       string `json:"subject"`
	Content       string `json:"content"`
	AttachedFiles string `json:"attached_file"`
	Type          string `json:"type"`
	Seen          *bool  `json:"seen"`
}
