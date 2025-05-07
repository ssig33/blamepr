# blamepr

A command-line tool that maps source code lines or files to the GitHub Pull Request (PR) that last modified them.

## Installation

### Using Go

```bash
# Install directly
go install github.com/ssig33/blamepr/cmd/blamepr@latest

# Or clone and build
git clone https://github.com/ssig33/blamepr.git
cd blamepr
make install
```

### Binary Downloads

Pre-compiled binaries for Linux, macOS, and Windows are available on the [releases page](https://github.com/ssig33/blamepr/releases).

## Usage

```bash
# Show the PR that last modified a file
blamepr path/to/file.go

# Show the PR that last modified a specific line in a file
blamepr path/to/file.go:123

# Open the PR in your default browser
blamepr -open path/to/file.go:123
```

## Authentication

`blamepr` requires GitHub API authentication to work properly. You can provide it in one of two ways:

1. Set the `GITHUB_TOKEN` environment variable with your GitHub personal access token:
   ```bash
   export GITHUB_TOKEN=your_github_token
   ```

2. Add your GitHub credentials to your `.netrc` file:
   ```
   machine github.com
   login your_username
   password your_github_token
   ```

## Requirements

- Git must be installed and available in your PATH
- The current directory must be within a Git repository
- The repository must be hosted on GitHub
- A GitHub API token with appropriate permissions

## License

WTFPL - Do What The Fuck You Want To Public License