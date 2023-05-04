package api

import (
	"errors"
	"os"
	"sinarmas/kredit-sinarmas/controllers/authentication"
	"sinarmas/kredit-sinarmas/controllers/kredit"

	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() error {
	origin_1 := os.Getenv("ORIGIN_1")
	if origin_1 == "" {
		return errors.New("origin setting not found")
	}

	s.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{origin_1},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Accept", "Content-Type", "Cookie", "Authorization"},
		AllowCredentials: true,
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
	s.Router.DELETE("/login", authHandler.IsAuthenticated, authHandler.Logout)
	s.Router.PATCH("/change_password", authHandler.IsAuthenticated, authHandler.ChangePassword)

	kreditRoutes := s.Router.Group("/kredit", authHandler.IsAuthenticated)
	{
		kreditRoutes.GET("/checklist_pencairan", kreditHandler.GetChecklistPencairan)
		kreditRoutes.PATCH("/checklist_pencairan", kreditHandler.UpdateChecklistPencairan)
		kreditRoutes.GET("/drawdown_report", kreditHandler.GetDrawdownReport)
	}

	return nil
}
