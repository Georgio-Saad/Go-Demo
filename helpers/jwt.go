package helpers

import (
	"time"
	"todogorest/constants"
	"todogorest/models"

	"github.com/kataras/jwt"
)

func GenerateAccessToken(user models.User) (token string, err error) {
	accessEnc, _, err := jwt.GCM(constants.AccessEncKey, nil)

	accessTokenClaims := &JWTClaims{
		User:      user,
		Claims:    jwt.Claims{Expiry: time.Now().Add(15 * time.Minute).Unix()},
		GrantType: constants.AccessToken,
	}

	accessToken, jwtErr := jwt.SignEncrypted(
		jwt.HS256,
		constants.AccessSignKey,
		accessEnc,
		accessTokenClaims,
		jwt.MaxAge(15*time.Minute),
	)

	return string(accessToken), jwtErr
}

func GenerateRefreshToken(user models.User) (token string, err error) {
	refreshEnc, _, refErr := jwt.GCM(constants.RefreshEncKey, nil)

	refreshTokenClaims := &JWTClaims{
		User:      user,
		Claims:    jwt.Claims{Expiry: time.Now().Add(2 * 7 * 24 * time.Hour).Unix()},
		GrantType: constants.RefreshToken,
	}

	refreshToken, refErr := jwt.SignEncrypted(
		jwt.HS256,
		constants.RefreshSignKey,
		refreshEnc,
		refreshTokenClaims,
		jwt.MaxAge(2*7*24*time.Hour),
	)

	return string(refreshToken), refErr
}
