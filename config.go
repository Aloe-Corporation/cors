package cors

import "time"

// Conf represents the configuration structure for CORS (Cross-Origin Resource
// Sharing) settings. It includes fields for specifying allowed origins, methods,
// headers, exposed headers, allowing credentials, and the maximum age of a CORS
// preflight request. It is actually a mirror of cors.Config structure.
type Conf struct {
	AllowOrigins     []string      `yaml:"allow_origins"`     // List of allowed origins. "*" means allows all origins.
	AllowMethods     []string      `yaml:"allow_methods"`     // List of allowed HTTP methods.
	AllowHeaders     []string      `yaml:"allow_headers"`     // List of allowed headers.
	ExposeHeaders    []string      `yaml:"expose_headers"`    // List of headers exposed to the browser.
	AllowCredentials bool          `yaml:"allow_credentials"` // Whether credentials can be included.
	MaxAge           time.Duration `yaml:"max_age"`           // Maximum age of a preflight request.
}
