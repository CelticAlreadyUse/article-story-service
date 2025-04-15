package usecase

import (
	"context"

	"github.com/CelticAlreadyUse/article-story-service/internal/model"
	"github.com/sirupsen/logrus"
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
func (u *CategoryUsecase)GetAll(ctx context.Context,params model.CategoryParams)([]*model.Category,error){
	logrus.WithFields(logrus.Fields{
		"params":params,
	})
	categoreis,err := u.repository.GetAll(ctx,params)
	if err !=nil{
		return nil,err
	}
	return categoreis,nil
}
func (u *CategoryUsecase)GetAllCategoriesByIds(ctx context.Context,id []int64)([]*model.Category,error){
	logrus.WithFields(logrus.Fields{
		"id_list":id,
	})
	Categories,err := u.repository.GetAllCategoriesByIds(ctx,id)
	if err !=nil{
		return nil,err
	}
	return Categories,nil
}
