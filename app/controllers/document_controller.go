package controllers

import (
	"docryze-backend/app/services"

	"github.com/gin-gonic/gin"
)

type DocumentController struct {
	documentService *services.DocumentService
}

func (c *DocumentController) CreateDocument(ctx *gin.Context) {
	c.documentService.Create("title", "subTitle", "content")
}
