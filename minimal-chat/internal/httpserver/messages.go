package httpserver

import (
	"minimal_chat/internal/models"
	"minimal_chat/internal/session"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListRoomMessages(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id64 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "invalid room id",
		})
		return
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	msgs, err := models.ListGroupMessages(uint(id64), limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "list messages failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": msgs,
	})
}

func ListPrivateMessages(c *gin.Context) {
	uidStr, _, ok := session.CurrentUser(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "unauthorized",
		})
		return
	}
	selfID64, _ := strconv.ParseUint(uidStr, 10, 64)

	peerStr := c.Query("peer")
	peerID64, err := strconv.ParseUint(peerStr, 10, 64)
	if err != nil || peerID64 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "invalid peer uid",
		})
		return
	}
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))

	msgs, err := models.ListPrivateMessages(uint(selfID64), uint(peerID64), limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "list messages failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": msgs,
	})
}
