package home

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"qrgo/platform/database"
)

func Handler(db *database.PostgresStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		total, err := db.GetTotalTickets()
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.HTML(http.StatusOK, "home.tmpl", gin.H{
			"Total": total,
		})
	}
}
