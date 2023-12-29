package validators

import (
	"github.com/gin-gonic/gin"
)

func EmitError(c *gin.Context, err error, status_code int) {
	c.SecureJSON(status_code, map[string]any{
		"error": err.Error(),
	})
}
