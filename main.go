package main

import (
	"fmt"
	"os"

	"wikiroute/core"
	"wikiroute/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	var err error

	pagesMap, err := core.ReadPagesData("")
	if err != nil {
		panic(fmt.Sprintf("Failed to read pages data: %v", err))
	}
	fmt.Println("pagesMap read successfully", len(pagesMap.IDToTitle))

	pageLinksMap, err := core.ReadPageLinksData("")
	if err != nil {
		panic(fmt.Sprintf("Failed to read page links data: %v", err))
	}
	fmt.Println("pageLinksMap read successfully", len(pageLinksMap.PageLinksMap))

	e := echo.New()
	handlers.RegisterRoutes(e, pagesMap, pageLinksMap)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Start(":" + port)
}