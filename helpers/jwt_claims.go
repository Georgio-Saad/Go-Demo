package helpers

import (
	"todogorest/models"

	"github.com/kataras/jwt"
)

type JWTClaims struct {
	models.User
	jwt.Claims
	GrantType string `json:"grant_type"`
}
