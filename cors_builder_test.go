package cors

import (
	"testing"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/stretchr/testify/assert"
)

type BuilderNewTestData struct {
	name           string
	expectedResult *Builder
}

var BuilderNewTestCases = [...]BuilderNewTestData{
	{
		name: "Success case",
		expectedResult: &Builder{
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

func TestBuilderNew(t *testing.T) {
	for _, testCase := range BuilderNewTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder := &Builder{}
			assert.Equal(t, testCase.expectedResult, builder.New())
		})
	}
}

type BuilderNewFromConfigTestData struct {
	name           string
	conf           *Conf
	expectedResult *Builder
}

var BuilderNewFromConfigTestCases = [...]BuilderNewFromConfigTestData{
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
		expectedResult: &Builder{
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

func TestBuilderNewFromConfig(t *testing.T) {
	for _, testCase := range BuilderNewFromConfigTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder := &Builder{}
			assert.Equal(t, testCase.expectedResult, builder.NewFromConfig(testCase.conf))
		})
	}
}

type BuilderWithOriginsTestData struct {
	name           string
	newOrigins     []string
	expectedResult *Builder
}

var BuilderWithOriginsTestCases = [...]BuilderWithOriginsTestData{
	{
		name:       "Overwrite AllowOrigins with localhost",
		newOrigins: []string{"localhost"},
		expectedResult: &Builder{
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

func TestBuilderWithOrigins(t *testing.T) {
	for _, testCase := range BuilderWithOriginsTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder := &Builder{}
			assert.Equal(t, testCase.expectedResult, builder.New().WithOrigins(testCase.newOrigins...))
		})
	}
}

type BuilderWithMethodsTestData struct {
	name           string
	newMethods     []string
	expectedResult *Builder
}

var BuilderWithMethodsTestCases = [...]BuilderWithMethodsTestData{
	{
		name:       "Overwrite AllowMethods with localhost",
		newMethods: []string{"PATCH"},
		expectedResult: &Builder{
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

func TestBuilderWithMethods(t *testing.T) {
	for _, testCase := range BuilderWithMethodsTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder := &Builder{}
			assert.Equal(t, testCase.expectedResult, builder.New().WithMethods(testCase.newMethods...))
		})
	}
}

type BuilderWithHeadersTestData struct {
	name           string
	newHeaders     []string
	expectedResult *Builder
}

var BuilderWithHeadersTestCases = [...]BuilderWithHeadersTestData{
	{
		name:       "Overwrite AllowHeaders with localhost",
		newHeaders: []string{"custom"},
		expectedResult: &Builder{
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

func TestBuilderWithHeaders(t *testing.T) {
	for _, testCase := range BuilderWithHeadersTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder := &Builder{}
			assert.Equal(t, testCase.expectedResult, builder.New().WithHeaders(testCase.newHeaders...))
		})
	}
}

type BuilderWithCredentialsTestData struct {
	name           string
	newCredentials bool
	expectedResult *Builder
}

var BuilderWithCredentialsTestCases = [...]BuilderWithCredentialsTestData{
	{
		name:           "Overwrite AllowCredentials with localhost",
		newCredentials: false,
		expectedResult: &Builder{
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

func TestBuilderWithCredentials(t *testing.T) {
	for _, testCase := range BuilderWithCredentialsTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder := &Builder{}
			assert.Equal(t, testCase.expectedResult, builder.New().WithCredentials(testCase.newCredentials))
		})
	}
}

type BuilderBuildTestData struct {
	name           string
	expectedResult *cors.Config
}

var BuilderBuildTestCases = [...]BuilderBuildTestData{
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

func TestBuilderBuild(t *testing.T) {
	for _, testCase := range BuilderBuildTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			builder := &Builder{}
			assert.Equal(t, testCase.expectedResult, builder.New().Build())
		})
	}
}
