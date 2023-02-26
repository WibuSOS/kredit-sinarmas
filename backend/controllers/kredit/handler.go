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
	var req RequestChecklistPencairan
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	// page := c.Query("page")
	// if page == "" {
	// 	page = "1"
	// }
	// limit := c.Query("limit")
	// if limit == "" {
	// 	limit = "10"
	// }

	res, err := h.Service.GetChecklistPencairan(&req)
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

// func (h *Handler) IsAuthenticated(c *gin.Context) {
// 	fullToken := c.GetHeader("Authorization")
// 	trimmedToken := strings.TrimPrefix(fullToken, "Bearer ")
// 	token, err := jwt.Parse(trimmedToken, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Authentication: Can't verify token!")
// 		}
// 		return []byte(os.Getenv("JWT_SECRET")), nil
// 	})

// 	if err != nil {
// 		log.Println(err.Error())
// 		c.AbortWithStatusJSON(http.StatusProxyAuthRequired, gin.H{
// 			"code":    http.StatusProxyAuthRequired,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	log.Println("Authentication: Token verified!")

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok || !token.Valid {
// 		log.Println("Authentication: Can't validate token!")
// 		c.AbortWithStatusJSON(http.StatusProxyAuthRequired, gin.H{
// 			"code":    http.StatusProxyAuthRequired,
// 			"message": "Authentication: Can't validate token!",
// 		})
// 		return
// 	}

// 	if err := claims.Valid(); err != nil {
// 		log.Println(err.Error())
// 		c.AbortWithStatusJSON(http.StatusProxyAuthRequired, gin.H{
// 			"code":    http.StatusProxyAuthRequired,
// 			"message": "Authentication: Can't validate time based claims!",
// 		})
// 		return
// 	}

// 	c.Set("userID", claims["ID"])
// 	c.Set("username", claims["Username"])
// 	c.Set("name", claims["Name"])
// 	c.Next()
// }
