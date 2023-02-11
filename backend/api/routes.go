package api

import (
	"sinarmas/kredit-sinarmas/controllers/authentication"
	"sinarmas/kredit-sinarmas/controllers/stagingCustomer"

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

	scRepo := stagingCustomer.NewRepository(s.DB)
	scService := stagingCustomer.NewService(scRepo)
	scHandler := stagingCustomer.NewHandler(scService)

	s.Router.POST("/login", authHandler.Login)
	s.Router.GET("/", authHandler.IsAuthenticated, scHandler.ValidateAndMigrate)

	// langRoutes := s.Router.Group("/:lang")
	// {
	// 	// auth controller (login)
	// 	langRoutes.POST("/login")
	// }

	return nil
}
