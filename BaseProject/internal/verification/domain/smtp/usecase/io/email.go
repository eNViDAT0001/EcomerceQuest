package io

type EmailForm struct {
	Subject     string
	Content     string
	To          []string
	Cc          []string
	Bcc         []string
	AttachFiles []string
}
