# bsdoc

A Go CLI tool that uses GitHub Copilot to generate detailed Backstage catalog documentation for your projects.

## Features

- 🤖 **AI-Powered**: Leverages GitHub Copilot SDK to intelligently analyze your project
- 📝 **Backstage Integration**: Generates compliant `catalog-info.yaml` files
- 🔍 **Smart Analysis**: Automatically detects project type, language, and structure
- ⚡ **Fast**: Built with Go for high performance

## Installation

### Homebrew

```bash
brew tap grahambrooks/bsdoc https://github.com/grahambrooks/bsdoc
brew install bsdoc
```

To upgrade to the latest release:

```bash
brew update && brew upgrade bsdoc
```

### Go

```bash
go install github.com/grahambrooks/bsdoc/cmd/bsdoc@latest
```

### Pre-built binaries

Download a release archive for your OS/arch from
[the releases page](https://github.com/grahambrooks/bsdoc/releases) and place
the `bsdoc` binary on your `PATH`.

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

## Releases

Releases are produced automatically by
[`.github/workflows/release.yml`](.github/workflows/release.yml) on every push
to `main` (and on manual `workflow_dispatch`). The workflow:

1. Runs `go test ./...`.
2. Computes a [CalVer](https://calver.org/) version of the form `YYYY.M.D`,
   appending `.N` when more than one release is cut on the same day
   (e.g. `2026.5.8`, `2026.5.8.1`, `2026.5.8.2`).
3. Cross-compiles `bsdoc` for `linux/amd64`, `linux/arm64`, `darwin/amd64`, and
   `darwin/arm64`, stamping the version into the binary via
   `-ldflags "-X main.version=$VERSION"`.
4. Tags the commit `vYYYY.M.D[.N]`, publishes a GitHub Release with the
   tarballs and a `SHA256SUMS` file.
5. Regenerates [`Formula/bsdoc.rb`](Formula/bsdoc.rb) with the new version and
   per-platform SHA256s and commits it back to `main` with `[skip ci]`.

To cut a release with a specific version, run the workflow manually from the
Actions tab and provide an override (e.g. `2026.5.8.3`).
