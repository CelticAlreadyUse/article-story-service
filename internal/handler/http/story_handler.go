package http_handler

import (
	"fmt"
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

var Validate = validator.New()

func (handler *storyHandler) RegisterRoute(e *echo.Echo) {
	g := e.Group("/v1/story")
	g.GET("/:id", handler.GetStoryByID) 
	g.POST("", handler.CreateStory)   
	g.GET("", handler.GetStories)  
	g.DELETE("/:id", handler.DeleteStoryByID) 
	g.PUT("/:id", handler.UpdateStory)
}

// For story to user
func (handler *storyHandler) CreateStory(c echo.Context) error {
	var body *model.Story
	err := c.Bind(&body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to get body data ")
	}
	err = Validate.Struct(body)

	if len(body.Content) == 0{
		return echo.NewHTTPError(http.StatusBadRequest, "content is required")
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	story, err := handler.storyUsecase.Create(c.Request().Context(), *body)
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, Response{
		Data:    story,
		Message: "sucessfully created story",
	})
}
func (handler *storyHandler) GetStoryByID(c echo.Context) error {
	id := c.Param("id")
	results, err := handler.storyUsecase.GetStoryByID(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "story id not found")
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
	if stories == nil {
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
func (handler *storyHandler) DeleteStoryByID(c echo.Context) error {
	id := c.Param("id")
	privId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to get an id")
	}
	err = handler.storyUsecase.Delete(c.Request().Context(), privId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, Response{
		Message: "Successfully deleting story",
	})
}
func (handler *storyHandler) UpdateStory(c echo.Context) error {
	var body *model.Story
	id := c.Param("id")
	err := c.Bind(&body)
	fmt.Println(body)
	if err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, "failed to get body")
	}
	primiId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to get an id")
	}
	story,storyCount, err := handler.storyUsecase.Update(c.Request().Context(), primiId, *body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if storyCount == 0 {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, "0 data been updated")
	}
	message := fmt.Sprintf("sucessfully updated %v items", storyCount)
	return c.JSON(http.StatusOK, Response{
		Data: story,
		Message: message,
	})
}

// for story to account
func (handler *storyHandler) GetStoryByAccountID(c echo.Context) error {
	accountId := c.Param("id")
	idInt, err := strconv.Atoi(accountId)
	if err != nil {
		return echo.ErrInternalServerError
	}
	stories, err := handler.storyUsecase.GetStoriesByUserID(c.Request().Context(), int64(idInt))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if stories == nil {
		return echo.NewHTTPError(echo.ErrNotFound.Code, "there is no stories for this users.")
	}
	messageRes := fmt.Sprintf("Sucesfully get user %v story", idInt)
	return c.JSON(http.StatusOK, Response{
		Data:    stories,
		Message: messageRes,
	})
}
