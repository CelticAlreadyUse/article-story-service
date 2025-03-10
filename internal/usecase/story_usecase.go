package usecase

import "github.com/CelticAlreadyUse/article-story-service/internal/model"


type storyUsecase struct{
	storyRepository model.StoryRepository
}

func  InitStoryUsecase (repo model.StoryRepository) *storyUsecase{
	return &storyUsecase{storyRepository: repo}
}