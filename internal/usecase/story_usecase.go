package usecase

import (
	"context"
	"errors"
	"strconv"

	"github.com/CelticAlreadyUse/article-story-service/internal/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type storyUsecase struct {
	storyRepository    model.StoryRepository
	categoryRepository model.CategoriesRepository
}

func InitStoryUsecase(storyRepo model.StoryRepository, categoryRepo model.CategoriesRepository) model.StoryUsecase {
	return &storyUsecase{storyRepository: storyRepo, categoryRepository: categoryRepo}
}

func (u *storyUsecase) Create(ctx context.Context, body model.Story) (*model.Story, error) {
	logrus.WithFields(logrus.Fields{
		"data": body,
	})
	var bdTagsInt []int64
	for _, tag := range body.TagsID {
		tagInt, err := strconv.Atoi(tag)
		if err != nil {
			return nil, errors.New("id type wrong")
		}
		bdTagsInt = append(bdTagsInt, int64(tagInt))
	}
	_, err := u.categoryRepository.GetAllCategoriesByIds(ctx, bdTagsInt)
	if err != nil {
		return nil, errors.New("id category doesn't found")
	}
	story, err := u.storyRepository.Create(ctx, body)
	if err != nil {
		return nil, err
	}
	return story, nil
}
func (u *storyUsecase) Delete(ctx context.Context, id primitive.ObjectID) error {
	logrus.WithFields(logrus.Fields{
		"id": id,
	})
	err := u.storyRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	logrus.Infof("story deleted %v", id)
	return u.storyRepository.Delete(ctx, id)
}
func (u *storyUsecase) Update(ctx context.Context, id primitive.ObjectID, storyBody model.Story) (*model.Story, int64, error) {
	logrus.WithFields(logrus.Fields{
		"_id":  id,
		"body": storyBody,
	})
	story, amount, err := u.storyRepository.Update(ctx, id, storyBody)
	if err != nil {
		logrus.Error(err)
		return nil, 0, err
	}
	if amount == 0 {
		logrus.Warn("no data been updated")
		return nil, 0, errors.New("0 data been updated")
	}
	return story, amount, nil
}
func (u *storyUsecase) GetAll(ctx context.Context, params *model.SearchParams) ([]model.Story, string, error) {
	logrus.WithFields(logrus.Fields{
		"params": params,
	})
	stories, nextcsr, err := u.storyRepository.GetAll(ctx, params)
	if err != nil {
		logrus.Error(err.Error())
		return nil, "", err
	}
	tagIDSet := map[string]struct{}{}

	for _, story := range stories {
		for _, tagID := range story.TagsID {
			tagIDSet[tagID] = struct{}{}
		}
	}

	var tagIDs []int64
	for id := range tagIDSet {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return nil, "", err
		}
		tagIDs = append(tagIDs, int64(idInt))
	}
	categories, err := u.categoryRepository.GetAllCategoriesByIds(ctx, tagIDs)
	if err != nil {
		return nil, "", err
	}
	catMap := make(map[int64]*model.Category)
	for _, cat := range categories {
		catMap[cat.ID] = cat
	}
	for i := range stories {
		for _, id := range stories[i].TagsID {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return nil, "", err
			}
			if cat, ok := catMap[int64(idInt)]; ok {
				stories[i].Tags = append(stories[i].Tags, cat)
			}
		}
	}
	return stories, nextcsr, nil
}
func (u *storyUsecase) GetStoryByID(ctx context.Context, userID string) (*model.Story, error) {
	logrus.WithFields(logrus.Fields{
		"id": userID,
	})
	oId, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	story, err := u.storyRepository.GetByID(ctx, oId)
	if err != nil {
		return nil, err
	}
	var bdTagsInt []int64
	for _, tag := range story.TagsID {
		tagInt, err := strconv.Atoi(tag)
		if err != nil {
			return nil, errors.New("id type wrong")
		}
		bdTagsInt = append(bdTagsInt, int64(tagInt))
	}
	category, err := u.categoryRepository.GetAllCategoriesByIds(ctx, bdTagsInt)
	if err != nil {
		return nil, err
	}
	story.TagsID = nil
	story.Tags = category
	return story, nil
}

func (u *storyUsecase) GetStoriesByUserID(ctx context.Context, id int64) ([]*model.Story, error) {
	logrus.WithFields(logrus.Fields{
		"user_id": id,
	})
	stories, err := u.storyRepository.GetStoriesByUserID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return stories, nil
}
