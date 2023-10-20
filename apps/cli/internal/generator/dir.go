package generator

import (
	"context"
	"errors"
	"os"

	"log/slog"
)

func (g *Generator) makeDirectories(_ context.Context) error {
	if err := os.MkdirAll(g.projectPath, os.ModePerm); err != nil {
		return err
	}
	if err := os.MkdirAll(g.projectPath+"apps/", os.ModePerm); err != nil {
		return err
	}
	if err := os.MkdirAll(g.projectPath+"packages/", os.ModePerm); err != nil {
		return err
	}

	slog.Info("creating backend and frontend directories", "backend", g.backendPath, "frontend", g.frontendPath)
	if err := os.MkdirAll(g.projectPath+"apps/"+g.backendPath, os.ModePerm); err != nil {
		return err
	}

	if err := os.MkdirAll(g.projectPath+"apps/"+g.frontendPath, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func (g *Generator) checkIfProjectExists(_ context.Context) error {
	// check if project exists
	fi, err := os.Stat(g.projectPath)
	if err == nil {
		return errors.New("project already exists: " + fi.Name())
	}
	return nil
}
