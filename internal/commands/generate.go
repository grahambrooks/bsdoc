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

func run(_ *cobra.Command, _ []string) error {

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
	defer func() {
		if err := session.Destroy(); err != nil {
			log.Printf("[WARN] failed to destroy session: %v", err)
		}
	}()

	// Event handler with improved logging
	session.On(func(event copilot.SessionEvent) {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		switch event.Type {
		case "message":
			fmt.Printf("\n[%s] [MESSAGE] %v\n", timestamp, event.Data)
		case "error":
			fmt.Printf("\n[%s] [ERROR] %v\n", timestamp, event.Data)
		case "progress":
			fmt.Printf("\n[%s] [PROGRESS] %v\n", timestamp, event.Data)
		case "done":
			fmt.Printf("\n[%s] [DONE] %v\n", timestamp, event.Data)
		default:
			fmt.Printf("\n[%s] [EVENT: %s] %v\n", timestamp, event.Type, event.Data)
		}
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

	event, err := session.SendAndWait(copilot.MessageOptions{Prompt: prompt}, 5*60*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Session started with ID: %v\n", event)

	return nil
}
