package authentication

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Login(c *gin.Context) {
	var req DataRequest

	domain := os.Getenv("DOMAIN")
	if domain == "" {
		log.Println("domain setting not found")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": errors.New("domain setting not found"),
		})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	res, err := h.Service.Login(req)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.SetCookie("authorization", res.Token, 86400, "/", domain, false, true)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "successful",
		"data":    res,
	})
}

func (h *Handler) Logout(c *gin.Context) {
	// domain := os.Getenv("DOMAIN")
	// if domain == "" {
	// 	log.Println("domain setting not found")
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"code":    http.StatusInternalServerError,
	// 		"message": errors.New("domain setting not found"),
	// 	})
	// 	return
	// }

	c.SetCookie("authorization", "", 0, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "successful",
	})
}

func (h *Handler) IsAuthenticated(c *gin.Context) {
	fullToken, err := c.Cookie("authorization")
	log.Println(fullToken)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
		})
		return
	}

	trimmedToken := strings.TrimPrefix(fullToken, "Bearer ")
	token, err := jwt.Parse(trimmedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("authentication: can't verify token")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": err.Error(),
		})
		return
	}
	log.Println("authentication: token verified")

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		log.Println("authentication: can't validate token")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "authentication: can't validate token",
		})
		return
	}

	if err := claims.Valid(); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "authentication: can't validate time based claims",
		})
		return
	}
	log.Println(claims)
	log.Printf("userID:%T", claims["id"])

	c.Set("userID", claims["id"])
	c.Set("username", claims["username"])
	c.Set("name", claims["name"])
	c.Next()
}

func (h *Handler) ChangePassword(c *gin.Context) {
	var req RequestChangePassword
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	userID := c.GetFloat64("userID")
	log.Println("userID:", userID)
	err := h.Service.ChangePassword(&req, uint(userID))
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
