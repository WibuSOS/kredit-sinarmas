package api

import (
	"log"
	"os"
	"sinarmas/kredit-sinarmas/controllers/automatedService"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"gorm.io/gorm"
)

type server struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func MakeServer(db *gorm.DB) *server {
	s := &server{
		Router: gin.Default(),
		DB:     db,
	}

	return s
}

func (s *server) RunServer() {
	port := os.Getenv("PORT")

	if err := s.SetupRouter(); err != nil {
		log.Panicln(err.Error())
	}

	if err := s.RunJobs(); err != nil {
		log.Panicln(err.Error())
	}

	if err := s.Router.Run(":" + port); err != nil {
		log.Panicln(err.Error())
	}
}

func (s *server) RunJobs() error {
	scheduler := gocron.NewScheduler(time.Local)

	asRepo := automatedService.NewRepository(s.DB)
	var asService automatedService.Service = automatedService.NewService(asRepo)

	if _, err := scheduler.Every(30).Minutes().Do(asService.ValidateAndMigrate); err != nil {
		return err
	}

	if _, err := scheduler.Every(15).Minutes().Do(asService.GenerateSkalaAngsuran); err != nil {
		return err
	}

	scheduler.StartAsync()
	return nil
}
