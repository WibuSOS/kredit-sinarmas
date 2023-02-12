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

	authRepo := authentication.NewRepository(s.DB)
	authService := authentication.NewService(authRepo)
	authHandler := authentication.NewHandler(authService)

	s.Router.POST("/login", authHandler.Login)
	s.Router.GET("/", authHandler.IsAuthenticated)

	// langRoutes := s.Router.Group("/:lang")
	// {
	// 	// auth controller (login)
	// 	langRoutes.POST("/login")
	// }

	return nil
}
