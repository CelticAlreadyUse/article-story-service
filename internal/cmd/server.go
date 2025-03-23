package cmd

import (
	"net/http"

	"github.com/CelticAlreadyUse/article-story-service/internal/config"
	"github.com/CelticAlreadyUse/article-story-service/internal/database/mongodb"
	http_handler "github.com/CelticAlreadyUse/article-story-service/internal/handler/http"
	"github.com/CelticAlreadyUse/article-story-service/internal/repository"
	"github.com/CelticAlreadyUse/article-story-service/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

var startServerCmd = &cobra.Command{
	Use:   "httpsrv",
	Short: "httpsrv is a command to run http server",
	Run: func(cmd *cobra.Command, args []string) {
		dbConn := mongodb.Connect()
		storyRepository := repository.InitStoryStruct(dbConn)
		storyUsecase := usecase.InitStoryUsecase(storyRepository)
		e := echo.New()
		e.GET("/ping", func(c echo.Context) error {
			return c.String(http.StatusOK, "pong")
		})
		storyHandler := http_handler.InitStoryHandler(storyUsecase)
		storyHandler.RegisterRoute(e)
		e.Start(":" + config.PORT_HTTP())
	},
}

func init() {
	rootCmd.AddCommand(startServerCmd)
}
