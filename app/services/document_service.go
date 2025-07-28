package services

import (
	"docryze-backend/app/models"
	"docryze-backend/app/repositories"
)

type DocumentService struct {
	documentRepository *repositories.DocumentRepository
}

func (s *DocumentService) Create(title string, subTitle string, content string) {
	err := s.documentRepository.Create(&models.Document{
		Status:   1,
		Title:    title,
		SubTitle: subTitle,
		Content:  content,
	})
	if err != nil {
		panic(err)
	}
}
