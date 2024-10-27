package home

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"qrgo/platform/database"
)

func Handler(db *database.PostgresStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		total, err := db.GetTotalTickets()
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		msg := "Log in"
		url := "/login"
		if sessions.Default(ctx).Get("profile") != nil {
			msg = "Log out"
			url = "/logout"
		}

		ctx.HTML(http.StatusOK, "home.tmpl", gin.H{
			"Total":  total,
			"BtnMsg": msg,
			"BtnUrl":    url,
		})
	}
}
