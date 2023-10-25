package middlewares

import (
	"net/http"
	"strings"
	"todogorest/constants"
	"todogorest/data/response"
	"todogorest/helpers"

	"github.com/gin-gonic/gin"
	"github.com/kataras/jwt"
)

func IsAdmin(ctx *gin.Context) {
	headerToken := ctx.Request.Header.Get("Authorization")

	var token string

	if len(headerToken) == 0 || !strings.HasPrefix(headerToken, "Bearer ") {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{StatusCode: http.StatusUnauthorized, Code: helpers.Unauthenticated, Data: response.ErrorMessage{Message: "Unauthenticated"}})
		ctx.Abort()
		return
	}

	token = strings.TrimPrefix(headerToken, "Bearer ")

	_, accessDec, err := jwt.GCM(constants.AccessEncKey, nil)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{StatusCode: http.StatusUnauthorized, Code: helpers.Unauthenticated, Data: response.ErrorMessage{Message: err.Error()}})
		ctx.Abort()
		return
	}

	decToken, jwtErr := jwt.VerifyEncrypted(jwt.HS256, constants.AccessSignKey, accessDec, []byte(token))

	if jwtErr != nil {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{StatusCode: http.StatusUnauthorized, Code: helpers.Unauthenticated, Data: response.ErrorMessage{Message: jwtErr.Error()}})
		ctx.Abort()
		return
	}

	var claims helpers.JWTClaims

	claimsErr := decToken.Claims(&claims)

	if claimsErr != nil {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{StatusCode: http.StatusUnauthorized, Code: helpers.Unauthenticated, Data: response.ErrorMessage{Message: claimsErr.Error()}})
		ctx.Abort()
		return
	}

	if claims.GrantType != constants.AccessToken {
		ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{StatusCode: http.StatusUnauthorized, Code: helpers.Unauthenticated, Data: response.ErrorMessage{Message: "Unauthorized"}})
		ctx.Abort()
		return
	}

	ctx.Set("user_id", claims.User.ID)

	ctx.Next()
}
