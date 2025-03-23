package http_handler

import (
	"net/http"

	"github.com/CelticAlreadyUse/article-story-service/internal/model"
	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
)

type storyHandler struct {
	storyUsecase model.StoryUsecase
}

func InitStoryHandler(storyUsecase model.StoryUsecase) *storyHandler {
	return &storyHandler{storyUsecase: storyUsecase}
}

var validate = validator.New()

func (handler *storyHandler) RegisterRoute(e *echo.Echo) {
	g := e.Group("/v1/story")
	g.GET("/:id", handler.GetStoryByID)
	g.POST("/create", handler.CreateStory)
}

func (handler *storyHandler) CreateStory(c echo.Context) error {
	var body *model.Story
	err := c.Bind(&body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to get body data ")
	}
	err = validate.Struct(body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = handler.storyUsecase.Create(c.Request().Context(), *body)
	return c.JSON(http.StatusAccepted, Response{
		Message: "sucessfully created story",
	})
}

func (handler *storyHandler) GetStoryByID(c echo.Context) error {
	id := c.Param("id")
	results, err := handler.storyUsecase.GetStoryByID(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "story id not found")
	}
	return c.JSON(http.StatusAccepted, Response{
		Message: "success",
		Data:    results,
	})
}
