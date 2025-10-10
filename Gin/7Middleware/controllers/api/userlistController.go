package api

import (
	"github.com/gin-gonic/gin"
)

type UserListController struct {
}

func (con UserListController) ApiInit(c *gin.Context) {
	c.String(200, "api默认页")
}

func (con UserListController) UserList(c *gin.Context) {
	c.String(200, "api下的用户列表")
}
