package main

import (
	"fmt"
	"os"
	"pretest-privyid/middleware"
	"strconv"

	echoMid "github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

const DefaultPort = 8000

// HTTPServerMain main function for serving services over http
func (s *Service) HTTPServerMain() {

	e := echo.New()
	e.Use(middleware.Logger)
	e.Use(echoMid.Recover())
	e.Use(echoMid.CORS())

	wGroup := e.Group("/public/api/v1")
	s.CategoryHandler.MountCategory(wGroup)
	s.ProductHandler.MountProduct(wGroup)

	// set REST port
	var port uint16
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		portInt, err := strconv.Atoi(portEnv)
		if err != nil {
			port = DefaultPort
		} else {
			port = uint16(portInt)
		}
	} else {
		port = DefaultPort
	}

	listenerPort := fmt.Sprintf(":%d", port)
	e.Logger.Fatal(e.Start(listenerPort))
}
