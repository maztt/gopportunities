package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maztt/gopportunities/schemas"
)

func ListOpeningsHandler(context *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		reportError(context, http.StatusInternalServerError, "error listing openings")
		return
	}

	reportSuccess(context, "list-openings", openings)
}
