package cmd

import (
	"log"
	"github.com/CelticAlreadyUse/article-story-service/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:"article-story-service",
	Short: "api service for story api gateaway",
}

func init(){
	config.InitLoadWithViper()
}

func Execute(){
	if err := rootCmd.Execute();
	err !=nil{
		log.Fatal(err)
	}
}