package routes

import (
	"log"
	_"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {

	//セッションからデータを破棄する
	session := sessions.Default(c)
	log.Println("セッション取得")
	session.Clear()
	log.Println("クリア処理")
	session.Save()

}

func Login(c *gin.Context, UserId string) {

	//セッションにデータを格納する
	session := sessions.Default(c)
	session.Set("UserId", UserId)
	session.Save()
}
