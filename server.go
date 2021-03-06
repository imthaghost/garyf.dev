package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	// logger
	e.Use(middleware.Logger())
	// Stream recovery
	e.Use(middleware.Recover())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD},
	}))
	// Static file handler
	e.Static("/", "assets")
	// routes
	e.File("/", "pages/index.html")
	e.File("/index", "pages/index.html")
	e.File("/projects", "pages/projects.html")
	e.File("/resume", "pages/resume.html")
	e.File("/articles", "pages/articles.html")
	e.File("/test", "pages/test.html")
	e.GET("/pdf", func(c echo.Context) error {
		return c.Attachment("assets/pdf/resume.pdf", "GaryFrederickResume.pdf")
	})
	// Server
	e.Logger.Fatal(e.Start(":8080"))

}
