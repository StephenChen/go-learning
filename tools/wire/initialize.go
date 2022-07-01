//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"context"
	"github.com/google/wire"
	"wire/foobarbaz"
)

func initializeBaz(ctx context.Context) (foobarbaz.Baz, error) {
	wire.Build(foobarbaz.SuperSet)
	return foobarbaz.Baz{}, nil
}
