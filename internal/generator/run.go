package generator

import (
	"context"
	"log/slog"
	"os/exec"
)

func (g *Generator) Run(ctx context.Context) error {
	err := g.checkTools(ctx)
	if err != nil {
		return err
	}

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
	out, err := exec.CommandContext(ctx, "moon", "init", "--minimal", "--yes", g.projectName).CombinedOutput()
	if err != nil {
		slog.ErrorContext(ctx, "moon init failed", "err", err, "out", string(out))
		return err
	}

	return nil
}
