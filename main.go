package main

import (
	"log/slog"
	"os"

	"github.com/ngoldack/gloomsy/cmd"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	cmd.Execute()
}
