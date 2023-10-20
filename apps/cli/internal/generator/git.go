package generator

import (
	"context"
	"log/slog"
	"os/exec"
)

func (g *Generator) initGit(ctx context.Context) error {
	out, err := exec.CommandContext(ctx, "git", "init", g.projectPath).CombinedOutput()
	if err != nil {
		slog.ErrorContext(ctx, "git init failed", "err", err, "out", string(out))
		return err
	}
	return nil
}
