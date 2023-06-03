package io

type CreateRequestContact struct {
	ID            uint   `json:"id"`
	Subject       string `json:"subject"`
	Content       string `json:"content"`
	AttachedFile string `json:"attached_file"`
	Type          string `json:"type"`
	Seen          bool   `json:"seen"`
}
