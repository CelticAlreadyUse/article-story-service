package cmd

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/CelticAlreadyUse/article-story-service/internal/config"
	"github.com/CelticAlreadyUse/article-story-service/internal/database/mongodb"
	grpchandler "github.com/CelticAlreadyUse/article-story-service/internal/handler/grpc"
	http_handler "github.com/CelticAlreadyUse/article-story-service/internal/handler/http"
	"github.com/CelticAlreadyUse/article-story-service/internal/repository"
	"github.com/CelticAlreadyUse/article-story-service/internal/usecase"
	pb "github.com/CelticAlreadyUse/article-story-service/pb/service"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var startServerCmd = &cobra.Command{
	Use:   "httpsrv",
	Short: "httpsrv is a command to run http server",
	Run: func(cmd *cobra.Command, args []string) {
		dbConn, ctx, cancel := mongodb.Connect()
		defer cancel()
		defer dbConn.Client().Disconnect(ctx)
		storyRepository := repository.InitStoryStruct(dbConn)
		storyUsecase := usecase.InitStoryUsecase(storyRepository)
		wg := new(sync.WaitGroup)
		wg.Add(2)
		go func() {
			defer wg.Done()
			e := echo.New()
			e.GET("/ping", func(c echo.Context) error {
				return c.String(http.StatusOK, "pong")
			})
			storyHandler := http_handler.InitStoryHandler(storyUsecase)
			storyHandler.RegisterRoute(e)
			e.Start(":" + config.PORT_HTTP())
		}()
		go func() {
			defer wg.Done()
			grpcServer := grpc.NewServer()
			storyHandler := grpchandler.InitStoryGrpcHandler(storyUsecase)
			pb.RegisterStoryServiceServer(grpcServer, storyHandler)
			port := fmt.Sprintf(":%v", config.GRCPPORT())
			httpListerner, err := net.Listen("tcp", port)
			if err != nil {
				panic("grpc server stoped")
			}
			logrus.Infof("gRPC server listening on %v", port)
			if err := grpcServer.Serve(httpListerner); err != nil {
				log.Fatalf("failed to serve gRPC server : %v", err)
			}
		}()
		wg.Wait()

	},
}

func init() {
	rootCmd.AddCommand(startServerCmd)
}
