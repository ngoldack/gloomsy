package cmd

import (
	"log/slog"
	"os"

	"github.com/ngoldack/gloomsy/internal/generator"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runNew,
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// required flags
	newCmd.Flags().StringP("name", "n", "", "Name of the new project")

	// optional flags
	newCmd.Flags().String("frontend", "", "Name of the frontend")
	newCmd.Flags().String("backend", "", "Name of the backend")

	newCmd.Flags().Bool("git", false, "Add git to the project")
	newCmd.Flags().Bool("dagger", false, "Add dagger to the project")
	newCmd.Flags().Bool("pulumi", false, "Add pulumi to the project")
	newCmd.Flags().Bool("fly", false, "Add fly to the project")
}

func runNew(cmd *cobra.Command, args []string) {
	ctx := cmd.Context()
	options := make([]generator.GeneratorOption, 0)
	// required flags
	projectName, err := cmd.Flags().GetString("name")
	if err != nil {
		panic(err)
	}
	options = append(options, generator.WithProjectName(projectName))

	// optional flags
	projectPath, err := cmd.Flags().GetString("path")
	if err == nil {
		options = append(options, generator.WithProjectPath(projectPath))
	} else {
		// add default project name
		options = append(options, generator.WithProjectPath(projectName+"/"))
	}

	frontendName, err := cmd.Flags().GetString("frontend")
	if err != nil {
		options = append(options, generator.WithFrontendName(frontendName), generator.WithFrontendPath(frontendName+"/"))
	} else {
		// add default frontend name
		options = append(options, generator.WithFrontendName("web"), generator.WithFrontendPath("web/"))
	}

	backendName, err := cmd.Flags().GetString("backend")
	if err != nil {
		options = append(options, generator.WithBackendName(backendName), generator.WithBackendPath(backendName+"/"))
	} else {
		// add default backend name
		options = append(options, generator.WithBackendName("api"), generator.WithBackendPath("api/"))
	}

	// git
	if git, err := cmd.Flags().GetBool("git"); err == nil && git {
		options = append(options, generator.WithGit())
	}

	// dagger
	if dagger, err := cmd.Flags().GetBool("dagger"); err == nil && dagger {
		options = append(options, generator.WithDagger())
	}

	// pulumi
	if pulumi, err := cmd.Flags().GetBool("pulumi"); err == nil && pulumi {
		options = append(options, generator.WithPulumi())
	}

	// fly
	if fly, err := cmd.Flags().GetBool("fly"); err == nil && fly {
		options = append(options, generator.WithFly())
	}

	g, err := generator.New(options...)
	if err != nil {
		panic(err)
	}

	err = g.Run(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Project generation failed", "err", err)
		os.Exit(1)
	}

	slog.Info("Project generated successfully")
}
