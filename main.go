package main

import (
	"fmt"
	"net/http"

	"github.com/Flyewzz/group_preparation/api"
	"github.com/Flyewzz/group_preparation/api/handlers/middleware"
	"github.com/Flyewzz/group_preparation/router"
	"github.com/spf13/viper"
)

func main() {
	PrepareConfig()
	r := router.NewRouter()
	r.StrictSlash(true)
	HandlerData := PrepareHandlerData()
	api.ConfigureHandlers(r, HandlerData)
	c := router.CorsSetup()
	corsHandler := c.Handler(r)
	amw := middleware.NewAuthenticationMiddleware()
	amw.SetUpExcludedRoutes()
	r.Use(amw.Middleware)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":"+viper.GetString("port"), corsHandler)
}
