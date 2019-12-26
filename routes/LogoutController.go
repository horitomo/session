package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostLogout(c *gin.Context) {
	log.Println("ログアウト処理")
	Logout(c) // 同じパッケージ内のログアウト処理

	// ログインフォームに戻す
	c.HTML(http.StatusOK, "login", gin.H{
	"UserId":       "",
	"ErrorMessage": "",
	})
}
