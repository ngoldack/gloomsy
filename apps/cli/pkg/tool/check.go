package tool

import (
	"context"
	"os/exec"
)

func Check(ctx context.Context, name string, args ...string) ([]byte, error) {
	cmd, err := Prepare(ctx, name, args...)
	if err != nil {
		return nil, err
	}

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return out, nil
}

func Prepare(_ context.Context, name string, args ...string) (*exec.Cmd, error) {
	path, err := exec.LookPath(name)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(path, args...)

	return cmd, err
}
