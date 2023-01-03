package api

import (
	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() error {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
	}))

	s.Router.GET("/")
	// langRoutes := s.Router.Group("/:lang")
	// {
	// 	// auth controller (login)
	// 	langRoutes.POST("/login")
	// }

	return nil
}
