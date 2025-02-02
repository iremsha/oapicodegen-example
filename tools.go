//go:build tools
// +build tools

package tools

import (
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
)


// generating pets
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config gen_configs/pets/config.yaml api/v2/pets/schema.yaml

// generating shops
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config gen_configs/shops/config.yaml api/v2/shops/schema.yaml
