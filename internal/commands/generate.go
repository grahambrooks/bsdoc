package commands

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/github/copilot-sdk/go"
	"github.com/spf13/cobra"
)

var projectPath string

func NewGenerateCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate backstage documentation for a project",
		Long:  `Uses GitHub Copilot to analyze your project and generate detailed backstage documentation.`,
		RunE:  run,
	}

	cmd.Flags().StringVarP(&projectPath, "path", "p", "", "Project path to analyze (default: current directory)")

	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	if projectPath == "" {
		var err error
		projectPath, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get current directory: %w", err)
		}
	}
	err := os.Chdir(projectPath)
	if err != nil {
		return err
	}

	client := copilot.NewClient(&copilot.ClientOptions{})
	if err := client.Start(); err != nil {
		log.Fatal(err)
	}
	defer client.Stop()

	// Create session
	session, err := client.CreateSession(&copilot.SessionConfig{
		Model: "gpt-5",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer session.Destroy()

	// Event handler
	session.On(func(event copilot.SessionEvent) {
		fmt.Printf("\n%s: %v\n", event.Type, event)
	})

	prompt := `Generate a comprehensive Backstage catalog-info.yaml file that includes:

- Component metadata (name, description, type, lifecycle, owner)
- Appropriate tags based on the project
- API specifications if applicable
- Dependencies if detected
- Links to documentation, repository, etc.\n

Use the Backstage catalog format version 1.0.0.

Make intelligent decisions about component type (service, library, website, etc.) based on the project structure.

Create 'docs' directory for each project and write markdown documentation for the component. Make sure that the documentation is referenced by the appropriate backstage entity definition.`

	event, err := session.SendAndWait(copilot.MessageOptions{Prompt: prompt}, 5*60*time.Second0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Session started with ID: %v\n", event)

	return nil
}
