package http_handler

import (
	"net/http"
	"strconv"

	"github.com/CelticAlreadyUse/article-story-service/internal/model"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"

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
	g.POST("", handler.CreateStory)
	g.GET("", handler.GetStories)
	g.DELETE("/:id",handler.DeleteStoryByID)
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
func (handler *storyHandler) GetStories(c echo.Context) error {
	var Params model.SearchParams

	categoryParams := c.QueryParams()["tags"]
	if len(categoryParams) > 5 {
		return echo.NewHTTPError(http.StatusBadRequest, "Maximum of 5 tags allowed")
	}
	Params.Tags = []string(categoryParams)

	keyWordParm := c.QueryParam("keyword")
	if keyWordParm != "" {
		Params.Keywords = keyWordParm
	}
	limitParam := c.QueryParam("limit")
	if limitParam != "" {
		limitInt, err := strconv.Atoi(limitParam)
		if err != nil {
			return echo.ErrInternalServerError
		}
		Params.Limit = int64(limitInt)
	}
	cursorParam := c.QueryParam("cursor")
	if cursorParam != "" {
		Params.Cursor = cursorParam
	}
	stories, nextcsr, err := handler.storyUsecase.GetAll(c.Request().Context(), Params)
	if err != nil {
		return echo.ErrBadRequest
	}
	if stories == nil{
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, Response{
		Data:    stories,
		Message: "success",
		Metadata: map[string]any{
			"length":      len(stories),
			"next_cursor": nextcsr,
			"has_more":    len(stories) == int(Params.Limit),
		},
	})
}
func (handler *storyHandler) DeleteStoryByID(c echo.Context)error{
	id :=c.Param("id")
	privId,err := primitive.ObjectIDFromHex(id)
	if err !=nil{
		return echo.NewHTTPError(http.StatusBadRequest,"failed to get an id")
	}
	err = handler.storyUsecase.DeleteByID(c.Request().Context(),privId)	
	if err !=nil{
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK, Response{
		Message: "Successfully deleting story",
	})
}