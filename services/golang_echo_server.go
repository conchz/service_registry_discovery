package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderWWWAuthenticate,
			echo.HeaderAuthorization},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
		<!DOCTYPE html>
		<html>
		<head>
		    <meta charset="UTF-8">
		    <meta http-equiv="X-UA-Compatible" content="IE=edge, chrome=1">
		    <meta name="viewport" content="width=device-width, initial-scale=1">
		    <title>Golang Echo</title>
		</head>
		<body>
		    <span>Welcome to Echo!</span>
		</body>
		`)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
