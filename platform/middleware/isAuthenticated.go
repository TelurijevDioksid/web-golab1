package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsAuthenticated(ctx *gin.Context) {
	if sessions.Default(ctx).Get("profile") == nil {
		reqUrl := ctx.Request.URL.Path
		if strings.Contains(reqUrl, "/info/") {
			id := strings.Split(reqUrl, "/")[2]
			session := sessions.Default(ctx)
			session.Set("id", id)
			if err := session.Save(); err != nil {
				ctx.String(http.StatusInternalServerError, err.Error())
				return
			}
		}
		ctx.Redirect(http.StatusSeeOther, "/login")
	} else {
		ctx.Next()
	}
}
