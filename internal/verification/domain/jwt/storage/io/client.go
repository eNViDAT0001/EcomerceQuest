package io

import (
	"strconv"
	"time"
)

type Client interface {
	GetClientID() string
	GetClientKey() string
	GetClientName() string
	GetUserID() string
}

type client struct {
	UserID    uint32    `json:"user_id" gorm:"column:user_id"`
	ClientID  string    `json:"client_id" gorm:"column:client_id"`
	ClientKey string    `json:"client_key" gorm:"column:client_key"`
	Name      string    `json:"name"`
	Active    bool      `json:"active" gorm:"column:active"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
}

func (s client) GetClientID() string {
	return s.ClientID
}
func (s client) GetUserID() string {
	return strconv.Itoa(int(s.UserID))
}

func (s client) GetClientKey() string {
	return s.ClientKey
}

func (s client) GetClientName() string {
	return s.Name
}
