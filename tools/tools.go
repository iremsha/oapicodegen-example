//go:build tools
// +build tools

package tools

import _ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config config.yaml api/v1/api.yaml
