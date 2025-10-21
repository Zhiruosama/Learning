package httpserver

import (
	"minimal_chat/internal/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginReq struct {
	Username string `json:"username" binding:"required,min=2,max=32"`
}

// POST /login
func Login(c *gin.Context) {
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	session.SaveAuthSession(c, req.Username, req.Username)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"uid":      req.Username,
			"username": req.Username,
		},
	})
}

// GET /me
func Me(c *gin.Context) {
	uid, username, ok := session.CurrentUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "unauthorized",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"uid":      uid,
			"username": username,
		},
	})
}

// GET /logout
func Logout(c *gin.Context) {
	session.ClearAuthSession(c)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}
