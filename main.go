package main

import (
	"fmt"
	"net/http"

	"github.com/Flyewzz/group_preparation/handlers"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

func main() {
	PrepareConfig()
	r := NewRouter()
	HandlerData := PrepareHandlerData()
	handlers.ConfigureHandlers(r, HandlerData)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"www.google.ru",
		},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	corsHandler := c.Handler(r)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":"+viper.GetString("port"), corsHandler)
}
