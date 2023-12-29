package auth

import (
	"flowcraft/auth-api/v2/src/utils/validators"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func NewHandler(g *gin.Engine) {
	router := g.Group("/auth")
	controller := NewAuthController()

	router.POST("/login", func(c *gin.Context) {
		var login *LoginDTO = &LoginDTO{}
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			validators.EmitError(c, err, 401)
			return
		}
		if err = login.Decode(body); err != nil {
			validators.EmitError(c, err, 401)
			return
		}
		status, data := controller.Login(*login)
		c.SecureJSON(status, data)
	})
	router.GET("/validate", func(c *gin.Context) {
		fmt.Println("validate")
		token := c.GetHeader("Authorization")
		if token == "" {
			validators.EmitError(c, fmt.Errorf("authorization header is required"), 401)
			return
		}
		status, data := controller.Authenticate(token)
		fmt.Println(status, data)
		c.SecureJSON(status, data)
	})
}
