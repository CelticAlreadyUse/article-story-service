package usecase

import (
	"context"

	"github.com/CelticAlreadyUse/article-story-service/internal/model"
)


type CategoryUsecase struct{
	repository 	model.CategoriesRepository
}
func  InitCategoryUsecase(repo model.CategoriesRepository)model.CategoriesUsecases{
	return &CategoryUsecase{repository: repo}
}
func (u *CategoryUsecase)Create(ctx context.Context,body model.Category)(model.Category,error){
	categry,err := u.repository.Create(ctx,body)
	
	if err !=nil{
		return model.Category{},err
	}
	return categry,nil
}
func (u *CategoryUsecase)GetByID(ctx context.Context,id int64)(*model.Category,error){
	panic("implement me")
}
func (u *CategoryUsecase)Delete(ctx context.Context,id int64)error{
	return u.repository.Delete(ctx,id)
}
func (u *CategoryUsecase)Update(ctx context.Context, id int64,body model.Category)(*model.Category,error){
	return u.repository.Update(ctx,id,body)
}