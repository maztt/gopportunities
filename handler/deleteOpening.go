package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maztt/gopportunities/schemas"
)

func DeleteOpeningHandler(context *gin.Context) {
	id := context.Query("id")
	if id == "" {
		reportError(context, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	// Find (and then) delete
	if err := db.First(&opening, id).Error; err != nil {
		reportError(context, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}
	if err := db.Delete(&opening).Error; err != nil {
		reportError(context, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}
	reportSuccess(context, "delete-opening", opening)
}
