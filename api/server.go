package api

// Generates the gin backend from the openapi.yaml
// This is here so that we can specify the version in go.mod and no binary installation required
//go:generate go tool oapi-codegen -config ./oapi-codegen.yaml ./openapi.yaml

// Enforce Server implementing ServerInterface
// See https://go.dev/doc/faq#guarantee_satisfies_interface
var (
	_ ServerInterface = Server{}
)

// Server implements the oapi-codegen's ServerInterface
type Server struct{}

// NewAPIServer creates a Server
func NewAPIServer() Server {
	return Server{}
}
