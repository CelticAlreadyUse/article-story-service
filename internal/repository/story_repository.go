package repository

import "gorm.io/gorm"


type StoryRepository struct {
	db *gorm.DB
}

func  InitStoryStruct (db *gorm.DB) StoryRepository {
	return StoryRepository{db: db}
}

