package httpserver

import (
	"crypto/md5"
	"fmt"
	"minimal_chat/internal/models"
	"minimal_chat/internal/session"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type loginReq struct {
	Username string `json:"username" binding:"required,min=2,max=32"`
	Password string `json:"password" binding:"required,min=3,max=64"`
	AvatarId string `json:"avatar_id" binding:"omitempty,max=32"`
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

	//MD5
	md5pwd := fmt.Sprintf("%x", md5.Sum([]byte(req.Password)))

	u, err := models.FindUserByUsername(req.Username)
	if err != nil || u.ID == 0 {
		//用户不存在 创建用户
		u, err = models.CreateUser(req.Username, md5pwd, req.AvatarId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "create user failed",
			})
			return
		}
	} else {
		//校验密码
		if u.Password != md5pwd {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "invalidcredentials",
			})
			return
		}
	}

	// 保存会话 使用用户id作为uid
	session.SaveAuthSession(c, strconv.Itoa(int(u.ID)), u.Username)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"uid":       u.ID,
			"username":  u.Username,
			"avatar_id": u.AvatarId,
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
