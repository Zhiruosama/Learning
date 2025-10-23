package session

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	store := cookie.NewStore([]byte("change_me_very_secret"))
	store.Options(sessions.Options{
		Path:     "/",
		Domain:   "",
		MaxAge:   30 * 24 * 60 * 60,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	return sessions.Sessions("minimal_chat", store)
}

// 保存登录状态信息
func SaveAuthSession(c *gin.Context, uid string, username string) {
	s := sessions.Default(c)
	s.Set("uid", uid)
	s.Set("username", username)
	_ = s.Save()
}

// 清除当前登录状态信息
func ClearAuthSession(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	_ = s.Save()
}

// 查看当前登陆状态信息
func CurrentUser(c *gin.Context) (uid string, username string, ok bool) {
	s := sessions.Default(c)
	u := s.Get("uid")
	n := s.Get("username")
	if u == nil {
		return "", "", false
	}
	return u.(string), n.(string), true
}
