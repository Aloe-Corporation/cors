package cors

import (
	"testing"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/stretchr/testify/assert"
)

type corsBuilderNewTestData struct {
	name           string
	expectedResult *CORSBuilder
}

var corsBuilderNewTestCases = [...]corsBuilderNewTestData{
	{
		name: "Success case",
		expectedResult: &CORSBuilder{
			Config: &cors.Config{
				AllowOrigins:     []string{"*"},
				AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
				AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
				ExposeHeaders:    []string{"Content-Length", "Content-Type"},
				AllowCredentials: true,
				MaxAge:           12 * time.Hour,
			},
		},
	},
}

func TestCORSBuilderNew(t *testing.T) {
	for _, testCase := range corsBuilderNewTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder := &CORSBuilder{}
			assert.Equal(t, testCase.expectedResult, builder.New())
		})
	}
}

type corsBuilderNewFromConfigTestData struct {
	name           string
	conf           *Conf
	expectedResult *CORSBuilder
}

var corsBuilderNewFromConfigTestCases = [...]corsBuilderNewFromConfigTestData{
	{
		name: "Success case",
		conf: &Conf{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
			AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
			ExposeHeaders:    []string{"Content-Length", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		},
		expectedResult: &CORSBuilder{
			Config: &cors.Config{
				AllowOrigins:     []string{"*"},
				AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
				AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
				ExposeHeaders:    []string{"Content-Length", "Content-Type"},
				AllowCredentials: true,
				MaxAge:           12 * time.Hour,
			},
		},
	},
}

func TestCORSBuilderNewFromConfig(t *testing.T) {
	for _, testCase := range corsBuilderNewFromConfigTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder := &CORSBuilder{}
			assert.Equal(t, testCase.expectedResult, builder.NewFromConfig(testCase.conf))
		})
	}
}

type corsBuilderWithOriginsTestData struct {
	name           string
	newOrigins     []string
	expectedResult *CORSBuilder
}

var corsBuilderWithOriginsTestCases = [...]corsBuilderWithOriginsTestData{
	{
		name:       "Overwrite AllowOrigins with localhost",
		newOrigins: []string{"localhost"},
		expectedResult: &CORSBuilder{
			Config: &cors.Config{
				AllowOrigins:     []string{"localhost"},
				AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
				AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
				ExposeHeaders:    []string{"Content-Length", "Content-Type"},
				AllowCredentials: true,
				MaxAge:           12 * time.Hour,
			},
		},
	},
}

func TestCORSBuilderWithOrigins(t *testing.T) {
	for _, testCase := range corsBuilderWithOriginsTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder := &CORSBuilder{}
			assert.Equal(t, testCase.expectedResult, builder.New().WithOrigins(testCase.newOrigins...))
		})
	}
}

type corsBuilderWithMethodsTestData struct {
	name           string
	newMethods     []string
	expectedResult *CORSBuilder
}

var corsBuilderWithMethodsTestCases = [...]corsBuilderWithMethodsTestData{
	{
		name:       "Overwrite AllowMethods with localhost",
		newMethods: []string{"PATCH"},
		expectedResult: &CORSBuilder{
			Config: &cors.Config{
				AllowOrigins:     []string{"*"},
				AllowMethods:     []string{"PATCH"},
				AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
				ExposeHeaders:    []string{"Content-Length", "Content-Type"},
				AllowCredentials: true,
				MaxAge:           12 * time.Hour,
			},
		},
	},
}

func TestCORSBuilderWithMethods(t *testing.T) {
	for _, testCase := range corsBuilderWithMethodsTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder := &CORSBuilder{}
			assert.Equal(t, testCase.expectedResult, builder.New().WithMethods(testCase.newMethods...))
		})
	}
}

type corsBuilderWithHeadersTestData struct {
	name           string
	newHeaders     []string
	expectedResult *CORSBuilder
}

var corsBuilderWithHeadersTestCases = [...]corsBuilderWithHeadersTestData{
	{
		name:       "Overwrite AllowHeaders with localhost",
		newHeaders: []string{"custom"},
		expectedResult: &CORSBuilder{
			Config: &cors.Config{
				AllowOrigins:     []string{"*"},
				AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
				AllowHeaders:     []string{"custom"},
				ExposeHeaders:    []string{"Content-Length", "Content-Type"},
				AllowCredentials: true,
				MaxAge:           12 * time.Hour,
			},
		},
	},
}

func TestCORSBuilderWithHeaders(t *testing.T) {
	for _, testCase := range corsBuilderWithHeadersTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder := &CORSBuilder{}
			assert.Equal(t, testCase.expectedResult, builder.New().WithHeaders(testCase.newHeaders...))
		})
	}
}

type corsBuilderWithCredentialsTestData struct {
	name           string
	newCredentials bool
	expectedResult *CORSBuilder
}

var corsBuilderWithCredentialsTestCases = [...]corsBuilderWithCredentialsTestData{
	{
		name:           "Overwrite AllowCredentials with localhost",
		newCredentials: false,
		expectedResult: &CORSBuilder{
			Config: &cors.Config{
				AllowOrigins:     []string{"*"},
				AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
				AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
				ExposeHeaders:    []string{"Content-Length", "Content-Type"},
				AllowCredentials: false,
				MaxAge:           12 * time.Hour,
			},
		},
	},
}

func TestCORSBuilderWithCredentials(t *testing.T) {
	for _, testCase := range corsBuilderWithCredentialsTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder := &CORSBuilder{}
			assert.Equal(t, testCase.expectedResult, builder.New().WithCredentials(testCase.newCredentials))
		})
	}
}

type corsBuilderBuildTestData struct {
	name           string
	expectedResult *cors.Config
}

var corsBuilderBuildTestCases = [...]corsBuilderBuildTestData{
	{
		name: "Success case",
		expectedResult: &cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
			AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
			ExposeHeaders:    []string{"Content-Length", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		},
	},
}

func TestCORSBuilderBuild(t *testing.T) {
	for _, testCase := range corsBuilderBuildTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder := &CORSBuilder{}
			assert.Equal(t, testCase.expectedResult, builder.New().Build())
		})
	}
}
