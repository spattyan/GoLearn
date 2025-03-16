package handler

import (
	"GoLearn/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(context *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		sendError(context,http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(context, "list-openings", openings)
}