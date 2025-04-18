package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/CelticAlreadyUse/article-story-service/internal/model"
	"github.com/CelticAlreadyUse/article-story-service/internal/usecase"
	"github.com/CelticAlreadyUse/article-story-service/internal/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCategoryUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.MockCategoryRepo)
	uc := usecase.InitCategoryUsecase(mockRepo)

	ctx := context.TODO()
	input := model.Category{Name: "Tech"}
	expected := model.Category{
		ID:        1,
		Name:      "Tech",
		Slug:      "tech",
		CreatedAt: time.Now(),
	}

	mockRepo.On("Create", ctx, input).Return(expected, nil)

	result, err := uc.Create(ctx, input)

	assert.NoError(t, err)
	assert.Equal(t, expected.Name, result.Name)
	assert.Equal(t, expected.Slug, result.Slug)
	mockRepo.AssertExpectations(t)
}

func TestCategoryUsecase_GetAll(t *testing.T) {
	mockRepo := new(mocks.MockCategoryRepo)
	uc := usecase.InitCategoryUsecase(mockRepo)

	ctx := context.TODO()
	params := model.CategoryParams{Keyword: "tech"}
	expected := []*model.Category{
		{ID: 1, Name: "Tech", Slug: "tech"},
		{ID: 2, Name: "Tech News", Slug: "tech-news"},
	}

	mockRepo.On("GetAll", ctx, params).Return(expected, nil)

	result, err := uc.GetAll(ctx, params)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, expected[0].Name, result[0].Name)
	mockRepo.AssertExpectations(t)
}

func TestCategoryUsecase_Delete(t *testing.T) {
	mockRepo := new(mocks.MockCategoryRepo)
	uc := usecase.InitCategoryUsecase(mockRepo)

	ctx := context.TODO()
	id := int64(1)

	mockRepo.On("Delete", ctx, id).Return(nil)

	err := uc.Delete(ctx, id)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryUsecase_Update(t *testing.T) {
	mockRepo := new(mocks.MockCategoryRepo)
	uc := usecase.InitCategoryUsecase(mockRepo)

	ctx := context.TODO()
	id := int64(1)
	input := model.Category{Name: "Science"}
	expected := &model.Category{
		ID:        1,
		Name:      "Science",
		Slug:      "science",
		UpdatedAt: time.Now(),
	}

	mockRepo.On("Update", ctx, id, input).Return(expected, nil)

	result, err := uc.Update(ctx, id, input)

	assert.NoError(t, err)
	assert.Equal(t, expected.Name, result.Name)
	mockRepo.AssertExpectations(t)
}

func TestCategoryUsecase_GetAllCategoriesByIds(t *testing.T) {
	mockRepo := new(mocks.MockCategoryRepo)
	uc := usecase.InitCategoryUsecase(mockRepo)

	ctx := context.TODO()
	ids := []int64{1, 2}
	expected := []*model.Category{
		{ID: 1, Name: "Tech"},
		{ID: 2, Name: "Science"},
	}

	mockRepo.On("GetAllCategoriesByIds", ctx, ids).Return(expected, nil)

	result, err := uc.GetAllCategoriesByIds(ctx, ids)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}
