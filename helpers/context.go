package helpers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetAuthUserId(ctx *gin.Context) string {
	userId, _ := ctx.Get("user_id")

	return fmt.Sprint(userId)
}
