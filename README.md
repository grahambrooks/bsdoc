# bsdoc

A Go CLI tool that uses GitHub Copilot to generate detailed Backstage catalog documentation for your projects.

## Features

- 🤖 **AI-Powered**: Leverages GitHub Copilot SDK to intelligently analyze your project
- 📝 **Backstage Integration**: Generates compliant `catalog-info.yaml` files
- 🔍 **Smart Analysis**: Automatically detects project type, language, and structure
- ⚡ **Fast**: Built with Go for high performance

## Installation

```bash
go install github.com/grahambrooks/bsdoc/cmd/bsdoc@latest
```

## Usage

### Generate Backstage Documentation

Generate a `catalog-info.yaml` file for your current project:

```bash
bsdoc generate
```

Specify a custom output path:

```bash
bsdoc generate --output /path/to/catalog-info.yaml
```

Generate documentation for a specific project:

```bash
bsdoc generate --path /path/to/project
```

### Version

```bash
bsdoc version
```

## Development

### Build

```bash
go build -o bin/bsdoc ./cmd/bsdoc
```

### Run

```bash
go run ./cmd/bsdoc
```

### Test

```bash
go test ./...
```
