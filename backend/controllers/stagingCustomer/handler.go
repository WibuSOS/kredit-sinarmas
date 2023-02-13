package stagingCustomer

import (
	"log"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) ValidateAndMigrate() {
	_, err := h.Service.ValidateAndMigrate()
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{
		// 	"message": err.Error(),
		// })
		log.Println(err.Error())
		return
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "OK",
	// 	"data":    res,
	// })
	log.Println("ValidateAndMigrate success")
}
