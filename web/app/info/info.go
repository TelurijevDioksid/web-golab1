package info

import (
	"net/http"
	"qrgo/platform/database"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Handler(db *database.PostgresStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		profile := session.Get("profile")

		ticket, err := db.GetTicket(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
			return
		}

		ctx.HTML(http.StatusOK, "info.tmpl", gin.H{
			"LoggedInUser": profile.(map[string]interface{})["nickname"],
			"Oib":          ticket.Oib,
			"FirstName":    ticket.FirstName,
			"LastName":     ticket.LastName,
			"CreatedAt":    ticket.CreatedAt,
		})
	}
}
