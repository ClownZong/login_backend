package login_backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.POST("/login", login)
	router.Run(":8080")
}

func login(c *gin.Context) {
	type Login struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}
	var loginInfo Login
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if loginInfo.User != "admin" || loginInfo.Password != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}
