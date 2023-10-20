package generator

import (
	"context"
	"errors"
)

func (g *Generator) generateFrontend(_ context.Context) error {
	return errors.ErrUnsupported
}

func (g *Generator) generateBackend(_ context.Context) error {
	return errors.ErrUnsupported
}
