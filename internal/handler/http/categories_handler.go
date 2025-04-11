package http_handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/CelticAlreadyUse/article-story-service/internal/helper"
	"github.com/CelticAlreadyUse/article-story-service/internal/model"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	Categoryusecase model.CategoriesUsecases
}

func InitCateoryHandler(usecase model.CategoriesUsecases) *CategoryHandler {
	return &CategoryHandler{Categoryusecase: usecase}
}
func (handler *CategoryHandler) RegisterRoute(e *echo.Echo) {
	g := e.Group("/v1/category")
	g.GET("", handler.GetAll)
	g.POST("", handler.Create)
	g.PUT("/:id",handler.Update)
	g.DELETE("/:id",handler.Delete)
}
func (handler *CategoryHandler) GetAll(c echo.Context) error {
	var params model.CategoryParams

	alphParams := c.QueryParam("alph")
	if alphParams != "" {
		params.Alph = alphParams
	}
	paramLimit := c.QueryParam("limit")
	if paramLimit != "" {
		limitInt, err := strconv.Atoi(paramLimit)
		if err != nil {
			return echo.ErrInternalServerError
		}
		params.Limit = int64(limitInt)
	}
	keyWord := c.QueryParam("keyword")
	if keyWord != "" {
		params.Keyword = keyWord
	}

	fmt.Println(params)
	return c.JSON(http.StatusOK, Response{
		Message: "sucessfully get category list",
	})
}
func (handler *CategoryHandler) Create(c echo.Context) error {
	var body *model.Category
	err := c.Bind(&body)
	if err != nil {
		return echo.ErrBadRequest
	}
	err = Validate.Struct(body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	isBlank := helper.NotBlank(body.Name)
	if !isBlank {
		return echo.NewHTTPError(http.StatusBadRequest, "categories name can't be blank")

	}
	category, err := handler.Categoryusecase.Create(c.Request().Context(), *body)
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return echo.NewHTTPError(echo.ErrBadGateway.Code, "categories already exist")
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, Response{
		Data:    category,
		Message: "sucessfully create category",
	})
}
func (handler *CategoryHandler) Update(c echo.Context) error {
	id := c.Param("id")
	var body *model.Category
	err := c.Bind(&body)
	if err != nil {
		return echo.ErrBadRequest
	}
	err = Validate.Struct(body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return echo.ErrInternalServerError
	}
	category, err := handler.Categoryusecase.Update(c.Request().Context(), int64(idInt), *body)
	if err != nil {
		return echo.ErrBadRequest
	}
	return c.JSON(http.StatusOK, Response{
		Data:    category,
		Message: "sucessfully update category ",
	})
}
func (handler *CategoryHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	idInt,err := strconv.Atoi(id)
	if err != nil {
		return echo.ErrBadRequest
	}
	err = handler.Categoryusecase.Delete(c.Request().Context(),int64(idInt))
	if err !=nil{
		return echo.NewHTTPError(http.StatusBadRequest,"Failed to delete category")
	}
	return c.JSON(http.StatusOK, Response{
		Message: "sucessfully deleting category",
	})
}
