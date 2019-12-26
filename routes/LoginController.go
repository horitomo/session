package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLogin(c *gin.Context) {

	c.HTML(http.StatusOK, "login", gin.H{
	"UserId":       "",
	"ErrorMessage": "",
})

}

func PostLogin(c *gin.Context) {
	log.Println("ログイン処理")
	UserId := c.PostForm("userId")

	Login(c, UserId) // // 同じパッケージ内のログイン処理

	c.Redirect(http.StatusMovedPermanently, "/menu/top")

}
