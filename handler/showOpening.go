package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maztt/gopportunities/schemas"
)

func ShowOpeningHandler(context *gin.Context) {
	id := context.Query("id")
	if id == "" {
		reportError(context, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		reportError(context, http.StatusNotFound, "opening not found")
		return
	}
	reportSuccess(context, "show-opening", opening)
}
