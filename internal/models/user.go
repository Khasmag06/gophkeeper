package models

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type User struct {
	Id        int       `json:"id,omitempty"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type TokenClaims struct {
	jwt.RegisteredClaims
	UserID string
}

type TokensResponse struct {
	AccessToken string `json:"accessToken"`
}
