package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetMenu(c *gin.Context) {
	UserId, _ := c.Get("UserId") // ログインユーザの取得

	c.HTML(http.StatusOK, "menu", gin.H{"UserId": UserId})
}
