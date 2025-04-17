// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/CelticAlreadyUse/article-story-service/internal/model"
	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// StoryRepository is an autogenerated mock type for the StoryRepository type
type StoryRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, story
func (_m *StoryRepository) Create(ctx context.Context, story model.Story) (*model.Story, error) {
	ret := _m.Called(ctx, story)

	var r0 *model.Story
	if rf, ok := ret.Get(0).(func(context.Context, model.Story) *model.Story); ok {
		r0 = rf(ctx, story)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Story)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Story) error); ok {
		r1 = rf(ctx, story)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *StoryRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx, params
func (_m *StoryRepository) GetAll(ctx context.Context, params model.SearchParams) ([]model.Story, string, error) {
	ret := _m.Called(ctx, params)

	var r0 []model.Story
	if rf, ok := ret.Get(0).(func(context.Context, model.SearchParams) []model.Story); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Story)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, model.SearchParams) string); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, model.SearchParams) error); ok {
		r2 = rf(ctx, params)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *StoryRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*model.Story, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Story
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) *model.Story); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Story)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStoriesByUserID provides a mock function with given fields: ctx, id
func (_m *StoryRepository) GetStoriesByUserID(ctx context.Context, id int64) ([]*model.Story, error) {
	ret := _m.Called(ctx, id)

	var r0 []*model.Story
	if rf, ok := ret.Get(0).(func(context.Context, int64) []*model.Story); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Story)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, story
func (_m *StoryRepository) Update(ctx context.Context, id primitive.ObjectID, story model.Story) (*model.Story, int64, error) {
	ret := _m.Called(ctx, id, story)

	var r0 *model.Story
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, model.Story) *model.Story); ok {
		r0 = rf(ctx, id, story)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Story)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID, model.Story) int64); ok {
		r1 = rf(ctx, id, story)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, primitive.ObjectID, model.Story) error); ok {
		r2 = rf(ctx, id, story)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewStoryRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewStoryRepository creates a new instance of StoryRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStoryRepository(t mockConstructorTestingTNewStoryRepository) *StoryRepository {
	mock := &StoryRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
