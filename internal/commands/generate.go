package commands

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
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
	logEvent("INFO", "Starting generate command")

	if projectPath == "" {
		var err error
		projectPath, err = os.Getwd()
		if err != nil {
			logEvent("ERROR", fmt.Sprintf("Failed to get current directory: %v", err))
			return fmt.Errorf("failed to get current directory: %w", err)
		}
	}
	err := os.Chdir(projectPath)
	if err != nil {
		logEvent("ERROR", fmt.Sprintf("Failed to change directory: %v", err))
		return err
	}

	ctx := context.Background()

	client := copilot.NewClient(&copilot.ClientOptions{})
	if err := client.Start(ctx); err != nil {
		logEvent("ERROR", fmt.Sprintf("Failed to start Copilot client: %v", err))
		log.Fatal(err)
	}
	defer client.Stop()

	logEvent("SESSION", "Creating Copilot session")
	session, err := client.CreateSession(ctx, &copilot.SessionConfig{
		Model: "gpt-5",
	})
	if err != nil {
		logEvent("ERROR", fmt.Sprintf("Failed to create Copilot session: %v", err))
		log.Fatal(err)
	}
	defer func() {
		if err := session.Disconnect(); err != nil {
			logEvent("WARN", fmt.Sprintf("Failed to disconnect session: %v", err))
		}
		logEvent("SESSION", "Session disconnected")
	}()

	resultChan := make(chan string, 1)
	errChan := make(chan error, 1)
	var content strings.Builder

	session.On(func(event copilot.SessionEvent) {
		ts := time.Now().Format("15:04:05")
		switch d := event.Data.(type) {
		case *copilot.AssistantMessageDeltaData:
			content.WriteString(d.DeltaContent)
			logEventWithTS(ts, "ASSISTANT_DELTA", d.DeltaContent)
		case *copilot.AssistantTurnEndData:
			logEventWithTS(ts, "ASSISTANT_TURN_END", "Assistant turn ended, result ready")
			resultChan <- content.String()
		case *copilot.SessionErrorData:
			logEventWithTS(ts, "SESSION_ERROR", d.Message)
			errChan <- fmt.Errorf("session error: %s", d.Message)
		case *copilot.UserMessageData:
			logEventWithTS(ts, "USER_MESSAGE", d.Content)
		case *copilot.ToolExecutionStartData:
			logEventWithTS(ts, "TOOL_EXEC_START", d.ToolName)
		case *copilot.ToolExecutionCompleteData:
			logEventWithTS(ts, "TOOL_EXEC_COMPLETE", d.ToolCallID)
		case *copilot.ToolExecutionProgressData:
			logEventWithTS(ts, "TOOL_EXEC_PROGRESS", d.ProgressMessage)
		case *copilot.AssistantMessageData:
			logEventWithTS(ts, "ASSISTANT_MESSAGE", d.Content)
		case *copilot.SessionUsageInfoData:
			logEventWithTS(ts, "SESSION_USAGE", fmt.Sprintf("currentTokens=%v messages=%v", d.CurrentTokens, d.MessagesLength))
		default:
			logEventWithTS(ts, "EVENT", string(event.Type()))
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

	logEvent("PROMPT", "Sending prompt to Copilot session")
	_, err = session.Send(ctx, copilot.MessageOptions{Prompt: prompt})
	if err != nil {
		logEvent("ERROR", fmt.Sprintf("Failed to send message: %v", err))
		return fmt.Errorf("failed to send message: %w", err)
	}

	select {
	case result := <-resultChan:
		logEvent("SUCCESS", "Documentation generation completed successfully")
		fmt.Printf("Result: %s\n", result)
		return nil
	case err := <-errChan:
		logEvent("FAILURE", fmt.Sprintf("Documentation generation failed: %v", err))
		return err
	case <-time.After(5 * 60 * time.Second):
		logEvent("TIMEOUT", "Timeout waiting for response from Copilot session")
		return fmt.Errorf("timeout waiting for response")
	}
}

func logEvent(eventType, message string) {
	ts := time.Now().Format("15:04:05")
	fmt.Printf("[%s] [%s] %s\n", ts, eventType, message)
}

func logEventWithTS(ts, eventType, message string) {
	fmt.Printf("[%s] [%s] %s\n", ts, eventType, message)
}
