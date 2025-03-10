package http_handler

import (
	"github.com/CelticAlreadyUse/article-story-service/internal/model"
	"github.com/labstack/echo/v4"
)

type storyHandler struct{
	storyUsecase model.StoryUsecase
}

func InitStoryHandler(storyUsecase model.StoryUsecase) *storyHandler{
 return &storyHandler{storyUsecase: storyUsecase}
}

func (handler storyHandler) RegisterRoute ( e *echo.Echo){

}	