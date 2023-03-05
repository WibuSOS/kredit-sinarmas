package api

import (
	"sinarmas/kredit-sinarmas/controllers/authentication"
	"sinarmas/kredit-sinarmas/controllers/kredit"

	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() error {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
	}))

	// authentication
	authRepo := authentication.NewRepository(s.DB)
	authService := authentication.NewService(authRepo)
	authHandler := authentication.NewHandler(authService)

	// kredit
	kreditRepo := kredit.NewRepository(s.DB)
	kreditService := kredit.NewService(kreditRepo)
	kreditHandler := kredit.NewHandler(kreditService)

	s.Router.POST("/login", authHandler.Login)

	kreditRoutes := s.Router.Group("/kredit", authHandler.IsAuthenticated)
	{
		kreditRoutes.GET("/checklist_pencairan", kreditHandler.GetChecklistPencairan)
		kreditRoutes.PATCH("/checklist_pencairan", kreditHandler.UpdateChecklistPencairan)
		kreditRoutes.GET("/drawdown_report", kreditHandler.GetDrawdownReport)
	}

	return nil
}
