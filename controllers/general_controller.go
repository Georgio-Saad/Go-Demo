package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"todogorest/data/response"
	"todogorest/helpers"

	"github.com/gin-gonic/gin"
)

type GeneralController struct{}

func (controller *GeneralController) GetGeneralData(ctx *gin.Context) {
	var generalData interface{}

	locale, _ := ctx.Get("locale")
	filepath := fmt.Sprintf("resources/%s.json", locale)

	fileData, _ := os.Open(filepath)

	defer fileData.Close()

	content, _ := io.ReadAll(fileData)

	err := json.Unmarshal(content, &generalData)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{StatusCode: http.StatusBadRequest, Code: helpers.BadRequest, Data: response.ErrorMessage{Message: err.Error()}})
		return
	}

	ctx.JSON(http.StatusOK, generalData)
}

func NewGeneralController() *GeneralController {
	return &GeneralController{}
}
