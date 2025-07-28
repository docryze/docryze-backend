package controllers

import "docryze-backend/app/services"

var (
	DocController *DocumentController
)

func init() {
	DocController = &DocumentController{
		documentService: services.DocService,
	}
}
