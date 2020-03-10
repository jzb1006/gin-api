package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID         int       `json:"id"`
	Openid     string    `json:"openid"`
	Unionid    string    `json:"unionid"`
	Nickname   string    `json:"nickname"`
	Gender     string    `json:"gender"`
	Language   string    `json:"language"`
	City       string    `json:"city"`
	Province   string    `json:"province"`
	Country    string    `json:"country"`
	AvatarUrl  string    `json:"avatar_url"`
	SessionKey string    `json:"session_key"`
}

