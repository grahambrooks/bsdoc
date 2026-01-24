package copilot

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	copilotSDK "github.com/github/copilot-sdk/go"
)

type DocGenerator struct {
	client *copilotSDK.Client
}

func NewDocGenerator() *DocGenerator {
	return &DocGenerator{
		client: copilotSDK.NewClient(nil),
	}
}

type ProjectInfo struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Language    string   `json:"language"`
	Files       []string `json:"files"`
	HasReadme   bool     `json:"has_readme"`
	ReadmeText  string   `json:"readme_text,omitempty"`
}

func (d *DocGenerator) GenerateBackstageDoc(ctx context.Context, projectPath string) (string, error) {
	if err := d.client.Start(); err != nil {
		return "", fmt.Errorf("failed to start copilot client: %w", err)
	}
	defer d.client.Stop()

	projectInfo, err := d.analyzeProject(projectPath)
	if err != nil {
		return "", fmt.Errorf("failed to analyze project: %w", err)
	}

	prompt := d.buildPrompt(projectInfo)

	session, err := d.client.CreateSession(&copilotSDK.SessionConfig{
		Model: "gpt-4o",
	})
	if err != nil {
		return "", fmt.Errorf("failed to create session: %w", err)
	}

	resultChan := make(chan string, 1)
	errChan := make(chan error, 1)
	var content strings.Builder

	session.On(func(event copilotSDK.SessionEvent) {
		if event.Type == "assistant.message" && event.Data.Content != nil {
			content.WriteString(*event.Data.Content)
		}
		if event.Type == "assistant.message.done" {
			resultChan <- content.String()
		}
		if event.Type == "error" {
			errChan <- fmt.Errorf("session error: %v", event)
		}
	})

	_, err = session.Send(copilotSDK.MessageOptions{Prompt: prompt})
	if err != nil {
		return "", fmt.Errorf("failed to send message: %w", err)
	}

	select {
	case result := <-resultChan:
		return result, nil
	case err := <-errChan:
		return "", err
	case <-time.After(60 * time.Second):
		return "", fmt.Errorf("timeout waiting for response")
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func (d *DocGenerator) analyzeProject(projectPath string) (*ProjectInfo, error) {
	info := &ProjectInfo{
		Name:  filepath.Base(projectPath),
		Files: []string{},
	}

	readmePath := filepath.Join(projectPath, "README.md")
	if data, err := os.ReadFile(readmePath); err == nil {
		info.HasReadme = true
		info.ReadmeText = string(data)
	}

	err := filepath.Walk(projectPath, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fileInfo.IsDir() {
			name := fileInfo.Name()
			if name == ".git" || name == "node_modules" || name == "vendor" || name == ".idea" {
				return filepath.SkipDir
			}
			return nil
		}

		relPath, _ := filepath.Rel(projectPath, path)
		info.Files = append(info.Files, relPath)

		if info.Language == "" {
			ext := filepath.Ext(path)
			info.Language = detectLanguage(ext)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return info, nil
}

func detectLanguage(ext string) string {
	languages := map[string]string{
		".go":   "Go",
		".js":   "JavaScript",
		".ts":   "TypeScript",
		".py":   "Python",
		".java": "Java",
		".rb":   "Ruby",
		".rs":   "Rust",
		".cs":   "C#",
		".cpp":  "C++",
		".c":    "C",
	}

	return languages[ext]
}

func (d *DocGenerator) buildPrompt(info *ProjectInfo) string {
	projectJSON, _ := json.MarshalIndent(info, "", "  ")

	var sb strings.Builder
	sb.WriteString("You are an expert at creating Backstage catalog-info.yaml files.\n\n")
	sb.WriteString("Given the following project information:\n\n")
	sb.WriteString("```json\n")
	sb.WriteString(string(projectJSON))
	sb.WriteString("\n```\n\n")

	if info.HasReadme {
		sb.WriteString("Project README.md content:\n\n")
		sb.WriteString("```markdown\n")
		sb.WriteString(info.ReadmeText)
		sb.WriteString("\n```\n\n")
	}

	sb.WriteString("Generate a comprehensive Backstage catalog-info.yaml file that includes:\n")
	sb.WriteString("- Component metadata (name, description, type, lifecycle, owner)\n")
	sb.WriteString("- Appropriate tags based on the project\n")
	sb.WriteString("- API specifications if applicable\n")
	sb.WriteString("- Dependencies if detected\n")
	sb.WriteString("- Links to documentation, repository, etc.\n\n")
	sb.WriteString("Use the Backstage catalog format version 1.0.0.\n")
	sb.WriteString("Make intelligent decisions about component type (service, library, website, etc.) based on the project structure.\n")
	sb.WriteString("For the owner field, use 'team/engineering' as a placeholder.\n\n")
	sb.WriteString("Return ONLY the YAML content, no explanations or markdown code blocks.")

	return sb.String()
}
