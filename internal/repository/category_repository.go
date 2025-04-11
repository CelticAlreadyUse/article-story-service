package repository

import (
	"context"
	"time"

	"github.com/CelticAlreadyUse/article-story-service/internal/helper"
	"github.com/CelticAlreadyUse/article-story-service/internal/model"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func InitCategoryRepository(db *gorm.DB) model.CategoriesRepository {
	return &categoryRepository{db: db}
}
func (r *categoryRepository) Create(ctx context.Context, body model.Category) (model.Category, error) {
	Category := &model.Category{
		Name:      body.Name,
		Slug:      helper.Slugify(body.Name),
		CreatedAt: time.Now()}
	results := r.db.Create(Category).WithContext(ctx)
	if results.Error != nil {
		return model.Category{}, results.Error
	}
	return *Category, nil
}
func (r *categoryRepository) GetByID(ctx context.Context, id int64) (*model.Category, error) {
	var category model.Category

	r.db.First(&category, id)
	if err := r.db.Error; err != nil {
		return nil, err
	}
	return &category, nil
}
func (r *categoryRepository) Delete(ctx context.Context, id int64) error {
	var Category *model.Category
	if err := r.db.Delete(&Category, id).Error; err != nil {
		return err
	}
	return nil
}
func (r *categoryRepository) Update(ctx context.Context, id int64, body model.Category) (*model.Category, error) {
	var Category *model.Category
	r.db.Model(&Category).Where("id", id).Updates(model.Category{Name: body.Name, Slug: helper.Slugify(body.Name), UpdatedAt: time.Now()})
	results, err := r.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return results, nil
}
