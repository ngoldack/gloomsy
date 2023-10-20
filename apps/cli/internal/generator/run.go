package generator

import (
	"context"
	"log/slog"
)

func (g *Generator) Run(ctx context.Context) error {
	tools, err := g.getTools(ctx)
	if err != nil {
		return err
	}
	slog.Info("All required tools found", "tools", tools)

	err = g.verifyToolVersions(ctx, tools)
	if err != nil {
		return err
	}
	slog.Info("All required tools have the correct version")

	err = g.checkIfProjectExists(ctx)
	if err != nil {
		return err
	}

	err = g.makeDirectories(ctx)
	if err != nil {
		return err
	}

	if g.git {
		err = g.initGit(ctx)
		if err != nil {
			return err
		}
	}

	// init moon project
	if err = g.generateMoonProject(ctx); err != nil {
		return err
	}

	// generate backend
	if err = g.generateBackend(ctx); err != nil {
		return err
	}

	// generate frontend
	if err = g.generateFrontend(ctx); err != nil {
		return err
	}

	return nil
}
