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
	g.GET("/user/:id",handler.GetStoryByUserID)
}

// For story to user
func (handler *storyHandler) CreateStory(c echo.Context) error {
	var body *model.Story
	err := c.Bind(&body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to get body data ")
	}
	err = Validate.Struct(body)

	if len(body.Content) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "content is required")
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	story, err := handler.storyUsecase.Create(c.Request().Context(), *body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrorResponse{
			Error:   err.Error(),
			Message: "failed to create story",
		})
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

	categoryParams := c.QueryParams()["t"]
	if len(categoryParams) > 5 {
		return echo.NewHTTPError(http.StatusBadRequest, "Maximum of 5 tags allowed")
	}
	Params.Tags = []string(categoryParams)

	keyWordParm := c.QueryParam("s")
	if keyWordParm != "" {
		Params.Keywords = keyWordParm
	}
	limitParam := c.QueryParam("l")
	if limitParam != "" {
		limitInt, err := strconv.Atoi(limitParam)
		if err != nil {
			return echo.ErrInternalServerError
		}
		Params.Limit = int64(limitInt)
	} else {
		Params.Limit = 10
	}
	cursorParam := c.QueryParam("c")
	if cursorParam != "" {
		Params.Cursor = cursorParam
	}
	stories, nextcsr, err := handler.storyUsecase.GetAll(c.Request().Context(), &Params)
	if err != nil {
		return echo.ErrBadRequest
	}
	if stories == nil {
		return echo.ErrNotFound
	}
	hasMore := false
	if len(stories) < int(Params.Limit) {
		hasMore = false
	} else {
		if nextcsr != "" {
			hasMore = true
		} else {
			hasMore = false
		}
	}
	return c.JSON(http.StatusOK, Response{
		Data:    stories,
		Message: "success",
		Metadata: map[string]any{
			"length":      len(stories),
			"next_cursor": nextcsr,
			"has_more":    hasMore,
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
	if err != nil {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, "failed to get body")
	}
	primiId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to get an id")
	}
	story, storyCount, err := handler.storyUsecase.Update(c.Request().Context(), primiId, *body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if storyCount == 0 {
		return echo.NewHTTPError(echo.ErrBadRequest.Code, "0 data been updated")
	}
	message := fmt.Sprintf("sucessfully updated %v items", storyCount)
	return c.JSON(http.StatusOK, Response{
		Data:    story,
		Message: message,
	})
}

// for story to account
func (handler *storyHandler) GetStoryByUserID(c echo.Context) error {
	accountId := c.Param("id")
	idInt, err := strconv.Atoi(accountId)
	cursor := c.QueryParam("c")
	limit := 8
	hasMore := false
	if err != nil {
		return echo.ErrInternalServerError
	}
	stories,nextCsr, err := handler.storyUsecase.GetStoriesByUserID(c.Request().Context(),int64(idInt),cursor)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if stories == nil {
		return echo.NewHTTPError(echo.ErrNotFound.Code, "there is no stories for this users.")
	}
	messageRes := fmt.Sprintf("Sucesfully get user %v story", idInt)
	if len(stories) < int(limit	) {
		hasMore = false
	} else {
		if nextCsr != "" {
			hasMore = true
		} else {
			hasMore = false
		}
	}
	return c.JSON(http.StatusOK, Response{
		Data:    stories,
		Metadata: map[string]any{
			"message": messageRes,
			"length":      len(stories),
			"next_cursor": nextCsr,
			"has_more":    hasMore,
		},
	})
}
