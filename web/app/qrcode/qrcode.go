package qrcode

import (
	"net/http"
	"os"
	"qrgo/platform/authenticator"
	"qrgo/platform/database"
	"qrgo/platform/models"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

func Handler(db *database.PostgresStorage, m2mauth *authenticator.M2MAuthenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dto := new(models.CreateTicketDto)
		if err := ctx.ShouldBindJSON(dto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// authHeader := ctx.GetHeader("Authorization")
		// if err := m2mauth.VerifyM2MToken(ctx.Request.Context(), authHeader); err != nil {
			// ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			// return
		// }

		if dto.Vatin == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "VATIN is required"})
			return
		}
		if dto.FirstName == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "First name is required"})
			return
		}
		if dto.LastName == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Last name is required"})
			return
		}

		count, err := db.GetTotalTicketsByOib(dto.Vatin)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if count >= 3 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "You have reached the limit of 3 tickets"})
			return
		}

		id, err := db.CreateTicket(dto)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		link := os.Getenv("BASE_URL") + "/info/" + id
		qrImage, err := qrGenImage(link)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.Data(http.StatusOK, "image/png", qrImage)
	}
}

func qrGenImage(s string) ([]byte, error) {
	png, err := qrcode.Encode(s, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return png, nil
}
