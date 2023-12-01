package cors

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type middlewareTestData struct {
	name               string
	headers            map[string][]string
	config             *cors.Config
	expectedStatusCode int
}

var middlewareTestCases = [...]middlewareTestData{
	{
		name:    "Success case",
		headers: make(map[string][]string),
		config: &cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET"},
			AllowHeaders:     []string{"Content-Type", "Content-Length"},
			ExposeHeaders:    []string{"Content-Length", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           1 * time.Hour,
		},
		expectedStatusCode: http.StatusOK,
	},
	{
		name: "Request rejected: invalid method",
		headers: map[string][]string{
			"Origin": {"http://localhost:8080"},
		},
		config: &cors.Config{
			AllowOrigins:     []string{"http://localhost:9090"},
			AllowMethods:     []string{"GET"},
			AllowHeaders:     []string{"Content-Type", "Content-Length"},
			ExposeHeaders:    []string{"Content-Length", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           1 * time.Hour,
		},
		expectedStatusCode: http.StatusForbidden,
	},
	{
		name:               "Success case with nil conf",
		config:             nil,
		expectedStatusCode: http.StatusOK,
	},
}

func TestMiddlware(t *testing.T) {
	for i, testCase := range middlewareTestCases {
		t.Run("Case"+strconv.Itoa(i), func(t *testing.T) {
			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)

			req := httptest.NewRequest("GET", "/", nil)
			req.Header = testCase.headers

			context.Request = req

			m := Middleware(testCase.config)
			m(context)

			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)
		})
	}
}
