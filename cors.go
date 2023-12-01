package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Middleware is a Gin middleware function that adds CORS (Cross-Origin
// Resource Sharing) support to the HTTP responses. It takes a cors.Config
// parameter and returns a Gin middleware handler. If the provided cors.Config
// is nil, it creates a default configuration using the CORSBuilder and applies
// it to the middleware.
//
// Use the CORSBuilder to build the cors.Config.
func Middleware(c *cors.Config) gin.HandlerFunc {
	if c == nil {
		builder := CORSBuilder{}
		c = builder.New().Build()
	}

	return cors.New(*c)
}
