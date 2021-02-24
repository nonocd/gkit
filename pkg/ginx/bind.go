package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Bind is binds the passed struct pointer
func Bind(s interface{}, c *gin.Context) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	if err := c.ShouldBindWith(s, b); err != nil {
		return err
	}
	return nil
}
