package cors

import (
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

var middlewareTestCases = []*Conf{
	DefaultConf(),
	{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Content-Type", "Content-Length"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           1 * time.Hour,
	},
	nil,
}

func TestMiddlware(t *testing.T) {
	for i, testCase := range middlewareTestCases {
		t.Run("Case"+strconv.Itoa(i), func(t *testing.T) {
			rr := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(rr)
			context.Request = httptest.NewRequest("GET", "/", nil)
			m := Middleware(testCase)
			m(context)
		})
	}
}
