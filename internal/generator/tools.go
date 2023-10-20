package generator

import (
	"context"
	"regexp"

	"log/slog"

	"github.com/ngoldack/gloomsy/pkg/tool"
)

func (g *Generator) checkTools(ctx context.Context) error {
	// check if moon is installed
	out, err := tool.Check(ctx, "moon", "--version")
	if err != nil {
		slog.ErrorContext(ctx, "moon not found", "err", err)
		return err
	}
	slog.InfoContext(ctx,
		"moon found",
		"version", regexp.MustCompile(`\d+\.\d+\.\d+`).FindString(string(out)),
		"version_raw", string(out),
	)

	// check if go is installed
	out, err = tool.Check(ctx, "go", "version")
	if err != nil {
		slog.ErrorContext(ctx, "go not found", "err", err)
		return err
	}
	slog.InfoContext(ctx,
		"go found",
		"version", regexp.MustCompile(`\d+\.\d+\.\d+`).FindString(string(out)),
		"version_raw", string(out),
	)

	// check if node is installed
	out, err = tool.Check(ctx, "node", "--version")
	if err != nil {
		slog.ErrorContext(ctx, "node not found", "err", err)
		return err
	}
	slog.InfoContext(ctx,
		"node found",
		"version", regexp.MustCompile(`\d+\.\d+\.\d+`).FindString(string(out)),
		"version_raw", string(out),
	)

	// check if pnpm is installed
	out, err = tool.Check(ctx, "pnpm", "--version")
	if err != nil {
		slog.ErrorContext(ctx, "pnpm not found", "err", err)
		return err
	}
	slog.InfoContext(ctx,
		"pnpm found",
		"version", regexp.MustCompile(`\d+\.\d+\.\d+`).FindString(string(out)),
		"version_raw", string(out),
	)

	if g.git {
		// check if git is installed
		out, err := tool.Check(ctx, "git", "--version")
		if err != nil {
			slog.ErrorContext(ctx, "git not found", "err", err)
			return err
		}
		slog.InfoContext(ctx,
			"git found",
			"version", regexp.MustCompile(`\d+\.\d+\.\d+`).FindString(string(out)),
			"version_raw", string(out),
		)
	}

	if g.dagger {
		// check if dagger is installed
		out, err := tool.Check(ctx, "dagger", "version")
		if err != nil {
			slog.ErrorContext(ctx, "dagger not found", "err", err)
			return err
		}
		slog.InfoContext(ctx,
			"dagger found",
			"version", regexp.MustCompile(`\d+\.\d+\.\d+`).FindString(string(out)),
			"version_raw", string(out),
		)
	}

	return nil
}
