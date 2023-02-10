package api

import (
	"sinarmas/kredit-sinarmas/controllers/authentication"

	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() error {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
	}))

	repo := authentication.NewRepository(s.DB)
	service := authentication.NewService(repo)
	handler := authentication.NewHandler(service)

	s.Router.GET("/login", handler.Login)

	// langRoutes := s.Router.Group("/:lang")
	// {
	// 	// auth controller (login)
	// 	langRoutes.POST("/login")
	// }

	return nil
}
