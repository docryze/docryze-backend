package services

import "docryze-backend/app/repositories"

var (
	DocService *DocumentService
)

func init() {
	DocService = &DocumentService{
		documentRepository: &repositories.DocumentRep,
	}
}
