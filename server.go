package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// logger
	e.Use(middleware.Logger())

	// Static file handler
	e.Static("/", "assets")

	// routes
	e.File("/", "pages/index.html")
	e.File("/index", "pages/index.html")
	e.File("/projects", "pages/projects.html")
	e.File("/articles", "pages/articles.html")

	// start server
	_ = e.Start(":8080")

}
