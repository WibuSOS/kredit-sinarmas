package api

import (
	"sinarmas/kredit-sinarmas/controllers/stagingCustomer"

	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() error {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
	}))

	repo := stagingCustomer.NewRepository(s.DB)
	service := stagingCustomer.NewService(repo)
	handler := stagingCustomer.NewHandler(service)

	s.Router.GET("/", handler.ValidateAndMigrate)

	// langRoutes := s.Router.Group("/:lang")
	// {
	// 	// auth controller (login)
	// 	langRoutes.POST("/login")
	// }

	return nil
}
