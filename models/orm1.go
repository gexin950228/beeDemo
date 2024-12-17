package models

import (
	"time"
)

type LoginUser struct {
	Id            int
	Username      string
	Password      string
	RedirectUri   string
	VerifyCode    string
	Address       string
	LastLoginTime time.Time
}
