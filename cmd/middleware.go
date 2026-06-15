package cmd

import (
	"e-wallet-transaction-service/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (d *Dependency) MiddlewareValidateToken(c *gin.Context) {

	var log = helpers.Logger

	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		log.Error("authorization empty")
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, "unauthorized", nil)
		c.Abort()
		return
	}

	tokenData, err := d.External.ValidateToken(c.Request.Context(), auth)
	if err != nil {
		log.Error("failed to validate token: ", err)
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, "unauthorized", nil)
		c.Abort()
	}

	c.Set("token", tokenData)

	c.Next()
}
