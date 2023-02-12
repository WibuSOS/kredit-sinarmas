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

// func (h *Handler) JoinRoom(c *gin.Context) {
// 	roomId := c.Param("room_id")
// 	userId := c.Param("user_id")
// 	langReq := c.Param("lang")
// 	localizator := c.MustGet("localizator")

// 	room, err := h.Service.JoinRoom(roomId, userId)
// 	if err != nil {
// 		errors.LogError(err)
// 		c.JSON(err.Status, gin.H{
// 			"message": localizator.(*language.Config).Lookup(langReq, err.Message),
// 		})
// 		return
// 	}

// 	statusArr := []string{"mulai transaksi", "barang dibayar", "barang dikirim", "konfirmasi barang sampai"}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message":  localizator.(*language.Config).Lookup(langReq, "successjoinroom"),
// 		"data":     room,
// 		"statuses": statusArr,
// 	})
// }

// func (h *Handler) JoinRoomPembeli(c *gin.Context) {
// 	roomId := c.Param("room_id")
// 	userId := c.Param("user_id")
// 	langReq := c.Param("lang")
// 	localizator := c.MustGet("localizator")
// 	err := h.Service.JoinRoomPembeli(roomId, userId)
// 	if err != nil {
// 		errors.LogError(err)
// 		c.JSON(err.Status, gin.H{
// 			"message": localizator.(*language.Config).Lookup(langReq, err.Message),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": localizator.(*language.Config).Lookup(langReq, "successjoinroombuyer"),
// 	})
// }
