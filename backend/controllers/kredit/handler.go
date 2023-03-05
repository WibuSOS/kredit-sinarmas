package kredit

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetChecklistPencairan(c *gin.Context) {
	page := c.Query("page")
	if page == "" {
		page = "1"
	}
	limit := c.Query("limit")
	if limit == "" {
		limit = "10"
	}

	res, err := h.Service.GetChecklistPencairan(page, limit)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "successful",
		"data":    res,
	})
}

func (h *Handler) UpdateChecklistPencairan(c *gin.Context) {
	var req RequestUpdateChecklistPencairan
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	log.Println("custcodes:", req.Custcodes)
	err := h.Service.UpdateChecklistPencairan(&req)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "successful",
	})
}
