package main

import (
	"context"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	if err = moonCI(client); err != nil {
		panic(err)
	}
}

func moonCI(client *dagger.Client) error {
	client = client.Pipeline("moon-ci")

	moon := client.Container().From("alpine:latest")
	moon = moon.WithExec([]string{"apk", "add", "--no-cache", "curl"})
	moon = moon.WithExec([]string{"curl", "-fsSL", "https://moonrepo.dev/install/moon.sh", "|", "bash"})

	src := client.Host().Directory("./")

	moon = moon.WithDirectory("/src", src).WithWorkdir("/src")
	moon.WithExec([]string{"moon", "ci"})

	return nil
}
