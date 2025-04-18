package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/CelticAlreadyUse/article-story-service/internal/model"
	"github.com/CelticAlreadyUse/article-story-service/internal/tests/mocks"
	"github.com/CelticAlreadyUse/article-story-service/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateStory_Success(t *testing.T) {
	storyRepo := new(mocks.MockStoryRepo)
	categoryRepo := new(mocks.MockCategoryRepo)
	uc := usecase.InitStoryUsecase(storyRepo, categoryRepo)

	text := "this is text test"
	ptrStr := &text
	dummyStory := model.Story{
		Title:   "My Story",
		TagsID:  []string{"1", "2"},
		Content: []*model.StoryElement{{Text: ptrStr, Type: "paragraph"}},
	}
	categoryRepo.On("GetAllCategoriesByIds", mock.Anything, dummyStory.TagsID).
		Return([]*model.Category{{ID: 1}, {ID: 2}}, nil)

	storyRepo.On("Create", mock.Anything, dummyStory).
		Return(&dummyStory, nil)

	result, err := uc.Create(context.Background(), dummyStory)

	assert.NoError(t, err)
	assert.Equal(t, "My Story", result.Title)

	storyRepo.AssertExpectations(t)
	categoryRepo.AssertExpectations(t)
}

func TestCreateStory_CategoryNotFound(t *testing.T) {
	storyRepo := new(mocks.MockStoryRepo)
	categoryRepo := new(mocks.MockCategoryRepo)
	uc := usecase.InitStoryUsecase(storyRepo, categoryRepo)
	text := "this is text test that should failed"
	ptrStr := &text
	dummyStory := model.Story{
		Title:   "Unknown Category Story",
		TagsID:  []string{"999"},
		Content: []*model.StoryElement{{Text: ptrStr,Type: "paragraph"}},
	}

	categoryRepo.On("GetAllCategoriesByIds", mock.Anything, dummyStory.TagsID).
		Return(nil, errors.New("not found"))

	result, err := uc.Create(context.Background(), dummyStory)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.EqualError(t, err, "id category doesn't found")
}
