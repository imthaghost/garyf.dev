package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()


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
