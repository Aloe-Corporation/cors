package cors

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Conf struct {
	AllowAllOrigins  bool          `yaml:"allow_all_origins"`
	AllowMethods     []string      `yaml:"allow_methods"`
	AllowHeaders     []string      `yaml:"allow_headers"`
	ExposeHeaders    []string      `yaml:"expose_headers"`
	AllowCredentials bool          `yaml:"allow_credentials"`
	MaxAge           time.Duration `yaml:"max_age"`
}

// DefaultConf provied a "open" CORS configuration. Be sure of what you want to do before use it in production.
func DefaultConf() *Conf {
	return &Conf{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
}

// Middleware which handle CORS policy.
func Middleware(c *Conf) gin.HandlerFunc {
	if c == nil {
		c = DefaultConf()
	}
	return cors.New(cors.Config{
		AllowAllOrigins:  c.AllowAllOrigins,
		AllowMethods:     c.AllowMethods,
		AllowHeaders:     c.AllowHeaders,
		ExposeHeaders:    c.ExposeHeaders,
		AllowCredentials: c.AllowCredentials,
		MaxAge:           c.MaxAge,
	})
}
