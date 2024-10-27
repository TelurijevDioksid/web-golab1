package router

import (
	"encoding/gob"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"qrgo/platform/authenticator"
	"qrgo/platform/database"
	"qrgo/platform/middleware"
	"qrgo/web/app/callback"
	"qrgo/web/app/home"
	"qrgo/web/app/info"
	"qrgo/web/app/login"
	"qrgo/web/app/logout"
	"qrgo/web/app/qrcode"
	"qrgo/web/app/user"
)

func New(m2mauth *authenticator.M2MAuthenticator, auth *authenticator.Authenticator, db *database.PostgresStorage) *gin.Engine {
	router := gin.Default()

	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	router.LoadHTMLGlob("web/template/*")

	router.GET("/home", home.Handler(db))
	router.POST("/gen", qrcode.Handler(db, m2mauth))
	router.GET("/info/:id", middleware.IsAuthenticated, info.Handler(db))

	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", middleware.IsAuthenticated, user.Handler)
	router.GET("/logout", logout.Handler)

	return router
}
