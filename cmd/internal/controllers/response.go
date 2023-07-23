package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, satusCode int, message string) {
	logrus.Error(message)
	c.JSON(satusCode, errorResponse{message})
}
