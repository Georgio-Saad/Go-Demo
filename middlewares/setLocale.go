package middlewares

import (
	"strings"
	"todogorest/constants"

	"github.com/gin-gonic/gin"
)

func SetLocale(ctx *gin.Context) {
	var locale string = constants.LocaleEN
	fetchedLocale := ctx.Request.Header.Get("Accept-Language")

	if len(fetchedLocale) > 0 && (strings.EqualFold(fetchedLocale, constants.LocaleAR) || strings.EqualFold(fetchedLocale, constants.LocaleEN) || strings.EqualFold(fetchedLocale, constants.LocaleFR) || strings.EqualFold(fetchedLocale, constants.LocaleIT)) {
		locale = strings.ToLower(fetchedLocale)
	}

	ctx.Set("locale", locale)
	ctx.Next()
}
