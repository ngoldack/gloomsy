package generator

import (
	"context"
	"errors"
	"regexp"

	"log/slog"

	"github.com/ngoldack/gloomsy/pkg/tool"
)

func (g *Generator) getTools(ctx context.Context) (map[string]string, error) {
	tools := make(map[string]string)

	// check if moon is installed
	out, err := tool.Check(ctx, "moon", "--version")
	if err != nil {
		slog.ErrorContext(ctx, "moon not found", "err", err)
		return nil, err
	}
	version := regexp.MustCompile(`\d+\.\d+\.\d+`).FindString(string(out))
	tools["moon"] = version
	slog.DebugContext(ctx,
		"moon found",
		"version", version,
		"version_raw", string(out),
	)

	// check if go is installed
	out, err = tool.Check(ctx, "go", "version")
	if err != nil {
		slog.ErrorContext(ctx, "go not found", "err", err)
		return nil, err
	}
	version = regexp.MustCompile(`\d+\.\d+\.\d+`).FindString(string(out))
	tools["go"] = version
	slog.DebugContext(ctx,
		"go found",
		"version", version,
		"version_raw", string(out),
	)

	// check if node is installed
	out, err = tool.Check(ctx, "node", "--version")
	if err != nil {
		slog.ErrorContext(ctx, "node not found", "err", err)
		return nil, err
	}
	version = regexp.MustCompile(`\d+\.\d+\.\d+`).FindString(string(out))
	tools["node"] = version
	slog.DebugContext(ctx,
		"node found",
		"version", version,
		"version_raw", string(out),
	)

	// check if pnpm is installed
	out, err = tool.Check(ctx, "pnpm", "--version")
	if err != nil {
		slog.ErrorContext(ctx, "pnpm not found", "err", err)
		return nil, err
	}
	version = regexp.MustCompile(`\d+\.\d+\.\d+`).FindString(string(out))
	tools["pnpm"] = version
	slog.DebugContext(ctx,
		"pnpm found",
		"version", version,
		"version_raw", string(out),
	)

	if g.git {
		// check if git is installed
		out, err := tool.Check(ctx, "git", "--version")
		if err != nil {
			slog.ErrorContext(ctx, "git not found", "err", err)
			return nil, err
		}
		version = regexp.MustCompile(`\d+\.\d+\.\d+`).FindString(string(out))
		tools["git"] = version
		slog.DebugContext(ctx,
			"git found",
			"version", version,
			"version_raw", string(out),
		)
	}

	if g.dagger {
		// check if dagger is installed
		out, err := tool.Check(ctx, "dagger", "version")
		if err != nil {
			slog.ErrorContext(ctx, "dagger not found", "err", err)
			return nil, err
		}
		version = regexp.MustCompile(`\d+\.\d+\.\d+`).FindString(string(out))
		tools["dagger"] = version
		slog.DebugContext(ctx,
			"dagger found",
			"version", version,
			"version_raw", string(out),
		)
	}

	return tools, nil
}

type outdatedTool struct {
	Name            string
	CurrentVersion  string
	RequiredVersion string
}

func (g *Generator) verifyToolVersions(ctx context.Context, tools map[string]string) error {

	outdated := make([]outdatedTool, 0)
	for k, v := range tools {
		slog.Debug("checking tool version", "tool", k, "version", v)
	}

	slog.WarnContext(ctx, "tool version verification not implemented yet")

	if len(outdated) > 0 {
		slog.ErrorContext(ctx, "some tools are outdated", "outdated_tools", outdated)
		var err error
		for _, v := range outdated {
			if err == nil {
				err = errors.New(v.Name + " is outdated")
				continue
			}

			err = errors.Join(err, errors.New(v.Name+" is outdated"))
		}
		return err
	}

	return nil
}
