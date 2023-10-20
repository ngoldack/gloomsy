package generator

import (
	"context"
	"log/slog"
	"os/exec"
	"strings"
)

func (g *Generator) generateMoonProject(ctx context.Context) error {
	out, err := exec.CommandContext(ctx, "moon", "init", "--minimal", "--yes", g.projectName).CombinedOutput()
	if err != nil {
		slog.ErrorContext(ctx, "moon init failed", "err", err, "out", string(out))
		return err
	}

	return nil
}

func (g *Generator) generateFrontend(ctx context.Context) error {
	slog.WarnContext(ctx, "frontend generation not implemented yet")
	return nil
}

func (g *Generator) generateBackend(ctx context.Context) error {
	module := g.projectName + "/apps/" + strings.TrimSuffix(g.backendPath, "/")

	slog.DebugContext(ctx, "generating backend", "module", module)

	cmd := exec.CommandContext(ctx, "go", "mod", "init", g.projectPath+"apps/"+strings.TrimSuffix(g.backendPath, "/"))
	cmd.Dir = g.projectPath + "apps/" + g.backendPath

	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.ErrorContext(ctx, "go mod init failed", "err", err, "out", string(out))
		return err
	}

	return nil
}
