package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	. "github.com/horitomo/session/SessionInfo"
	"github.com/horitomo/session/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var LoginInfo SessionInfo

func main() {
	router := gin.Default()

	//テンプレートの設定
	router.LoadHTMLGlob("**/view/*.html")

	// セッションの設定
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.GET("/login", routes.GetLogin)
	router.POST("/login", routes.PostLogin)

	menu := router.Group("/menu")
	menu.Use(sessionCheck())
	{
	menu.GET("/top", routes.GetMenu)
	}
	router.POST("/logout", routes.PostLogout)

	router.Run(":8080")

}

func sessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

	session := sessions.Default(c)
	LoginInfo.UserId = session.Get("UserId")

	// セッションがない場合、ログインフォームをだす
	if LoginInfo.UserId == nil {
		log.Println("ログインしていません")
		c.Redirect(http.StatusMovedPermanently, "/login")
		c.Abort() // これがないと続けて処理されてしまう
	} else {
		c.Set("UserId", LoginInfo.UserId) // ユーザidをセット
		c.Next()
	}
	log.Println("ログインチェック終わり")
	}
}
