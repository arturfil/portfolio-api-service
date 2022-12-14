package main

import (
	"fmt"

	"github.com/arturfil/portfolio-api-service/internal/project"
	transportHttp "github.com/arturfil/portfolio-api-service/internal/transport/http"
)

func Run() error {
	fmt.Println("Starting up the api server...")

	projectService := project.NewService()

	appHandler := transportHttp.NewHandler(
		projectService,
	)

	if err := appHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Api is working!")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
