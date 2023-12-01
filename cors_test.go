package cors

import (
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type middlewareTestData struct {
	name   string
	config *cors.Config
}

var middlewareTestCases = [...]middlewareTestData{
	{
		name: "Success case",
		config: &cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET"},
			AllowHeaders:     []string{"Content-Type", "Content-Length"},
			ExposeHeaders:    []string{"Content-Length", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           1 * time.Hour,
		},
	},
	{
		name:   "Success case with nil conf",
		config: nil,
	},
}

func TestMiddlware(t *testing.T) {
	for i, testCase := range middlewareTestCases {
		t.Run("Case"+strconv.Itoa(i), func(t *testing.T) {
			rr := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(rr)
			context.Request = httptest.NewRequest("GET", "/", nil)
			m := Middleware(testCase.config)
			m(context)
		})
	}
}
