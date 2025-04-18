package mocks

import (
    "context"

    "github.com/CelticAlreadyUse/article-story-service/internal/model"
    "github.com/stretchr/testify/mock"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type MockStoryRepo struct {
    mock.Mock
}

func (m *MockStoryRepo) Create(ctx context.Context, story model.Story) (*model.Story, error) {
    args := m.Called(ctx, story)
    return args.Get(0).(*model.Story), args.Error(1)
}

func (m *MockStoryRepo) Delete(ctx context.Context, id primitive.ObjectID) error {
    args := m.Called(ctx, id)
    return args.Error(0)
}

func (m *MockStoryRepo) GetByID(ctx context.Context, id primitive.ObjectID) (*model.Story, error) {
    args := m.Called(ctx, id)
    return args.Get(0).(*model.Story), args.Error(1)
}

func (m *MockStoryRepo) GetAll(ctx context.Context, params *model.SearchParams) ([]model.Story, string, error) {
    args := m.Called(ctx, params)
    return args.Get(0).([]model.Story), args.String(1), args.Error(2)
}

func (m *MockStoryRepo) Update(ctx context.Context, id primitive.ObjectID, story model.Story) (*model.Story, int64, error) {
    args := m.Called(ctx, id, story)
    return args.Get(0).(*model.Story), args.Get(1).(int64), args.Error(2)
}

func (m *MockStoryRepo) GetStoriesByUserID(ctx context.Context, id int64) ([]*model.Story, error) {
    args := m.Called(ctx, id)
    return args.Get(0).([]*model.Story), args.Error(1)
}

type MockCategoryRepo struct {
    mock.Mock
}
func (m *MockCategoryRepo) Create(ctx context.Context, body model.Category) (model.Category, error) {
	args := m.Called(ctx, body)
	return args.Get(0).(model.Category), args.Error(1)
}
func (m *MockCategoryRepo) GetByID(ctx context.Context, id int64) (*model.Category, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Category), args.Error(1)
}

func (m *MockCategoryRepo) Delete(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockCategoryRepo) Update(ctx context.Context, id int64, body model.Category) (*model.Category, error) {
	args := m.Called(ctx, id, body)
	return args.Get(0).(*model.Category), args.Error(1)
}

func (m *MockCategoryRepo) GetAll(ctx context.Context, params model.CategoryParams) ([]*model.Category, error) {
	args := m.Called(ctx, params)
	return args.Get(0).([]*model.Category), args.Error(1)
}

func (m *MockCategoryRepo) GetAllCategoriesByIds(ctx context.Context, ids []int64) ([]*model.Category, error) {
	args := m.Called(ctx, ids)
	return args.Get(0).([]*model.Category), args.Error(1)
}
