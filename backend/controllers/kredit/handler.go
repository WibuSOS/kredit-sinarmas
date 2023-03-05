package kredit

import (
	"log"
	"net/http"
	"time"

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

func (h *Handler) GetDrawdownReport(c *gin.Context) {
	page := c.Query("page")
	if page == "" {
		page = "1"
	}
	limit := c.Query("limit")
	if limit == "" {
		limit = "10"
	}
	company := c.Query("company")
	branch := c.Query("branch")
	startDate := c.Query("start_date")
	if _, err := time.Parse("2006-01-02", startDate); err != nil {
		startDate = ""
	}
	endDate := c.Query("end_date")
	if _, err := time.Parse("2006-01-02", endDate); err != nil {
		endDate = ""
	}
	approvalStatus := c.Query("approval_status")
	if approvalStatus != "1" && approvalStatus != "0" {
		approvalStatus = ""
	}

	res, err := h.Service.GetDrawdownReport(page, limit, company, branch, startDate, endDate, approvalStatus)
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
