package repositories

import (
	"docryze-backend/app/models"
)

type DocumentRepository struct {
}

func (r *DocumentRepository) Create(document *models.Document) error {
	return db.Create(document).Error
}

func (r *DocumentRepository) GetByID(id uint) (*models.Document, error) {
	var document models.Document
	err := db.First(&document, id).Error
	return &document, err
}
