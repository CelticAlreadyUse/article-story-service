package usecase

import (
	"context"

	"github.com/CelticAlreadyUse/article-story-service/internal/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type storyUsecase struct {
	storyRepository model.StoryRepository
}

func InitStoryUsecase(repo model.StoryRepository) model.StoryUsecase {
	return &storyUsecase{storyRepository: repo}
}

func (u *storyUsecase) Create(ctx context.Context, story model.Story) error {
	logrus.WithFields(logrus.Fields{
		"data": story,
	})
	return u.storyRepository.Create(ctx,story) 
}
func (u *storyUsecase) DeleteByID(ctx context.Context, id int64) error {
	panic("implement me!")
}
func (u *storyUsecase) UpdateByID(ctx context.Context, id int64, story model.Story) (model.Story, error) {
	panic("implement me")
}
func (u *storyUsecase) GetAll(ctx context.Context,params model.SearchParams) ([]model.Story,string,error) {
	logrus.WithFields(logrus.Fields{
		"params" : params,
	})
	stories,nextcsr,err := u.storyRepository.GetAll(ctx,params)
	if err !=nil{
		logrus.Error(err.Error())
		return nil,"",err
	}
	return stories,nextcsr,nil
}
func (u *storyUsecase) GetStoryByID(ctx context.Context, userID string) (*model.Story, error) {
	logrus.WithFields(logrus.Fields{
		"id" : userID,
	})
	oId,err := primitive.ObjectIDFromHex(userID)
	if err !=nil{
		return nil,err
	}
	story,err := u.storyRepository.GetByID(ctx,oId)
		if err !=nil{
			return nil,err
		}
	return story,nil
}
