package models

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	Status   int
	Title    string
	SubTitle string
	Content  string
}

func (Document) TableName() string {
	return "t_document"
}
