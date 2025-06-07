package models

import (
	"github.com/golang-jwt/jwt"
)

// JWTClaim define auth claim struct
type JWTClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// LoginRequest defines logs struct
type LoginRequest struct {
	Password string `json:"password" form:"password" validate:"required,min=1"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	//Csrf     string `json:"csrf" form:"csrf" validate:"required"`
}
