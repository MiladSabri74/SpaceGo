# File-Manager

A professional CLI tool for managing files, built with Go.

## Project Structure

This project follows Go best practices for project layout:

```
File-Manager/
├── cmd/
│   └── filemanager/      # Main application entry point
│       └── main.go
├── internal/
│   └── fileops/          # Private business logic (not importable by other projects)
│       ├── ops.go        # Core file operations implementation
│       └── ops_test.go   # Unit tests for file operations
├── pkg/
│   └── manager/          # Public API (importable by other projects)
│       ├── manager.go    # Manager wrapper with CLI-friendly methods
│       └── manager_test.go # Unit tests for manager package
├── go.mod                # Go module definition
├── go.sum                # Dependency checksums
└── README.md             # This file
```

### Architecture

- **`cmd/filemanager`**: Application entry point that handles CLI argument parsing and command routing
- **`internal/fileops`**: Core file operations (List, Copy) - private to this module
- **`pkg/manager`**: Public API that wraps internal operations with additional utilities

This separation allows:
- Clear boundaries between layers
- Easy testing of each component
- Potential reuse of `pkg/manager` in other projects
- Protection of internal implementation details

## Features

- **list**: Display all files and directories in a given path
- **copy**: Copy files from source to destination while preserving permissions
- **help**: Display usage information

## Installation

### Option 1: Build from Source (Local Development)

```bash
cd File-Manager
go build -o file-manager ./cmd/filemanager
```

The binary will be created in the current directory. Run it with:
```bash
./file-manager help
```

### Option 2: Install to GOPATH/bin (Global Installation)

```bash
# Navigate to the project directory
cd File-Manager

# Install the binary to $GOPATH/bin
go install ./cmd/filemanager

# Add GOPATH/bin to your PATH (if not already added)
# For bash/zsh, add this line to your ~/.bashrc or ~/.zshrc:
export PATH=$PATH:$(go env GOPATH)/bin

# Then reload your shell configuration:
source ~/.bashrc  # or source ~/.zshrc

# Now you can use file-manager from anywhere
file-manager help
```

**Note:** If you don't want to modify your PATH, you can run the binary directly from:
```bash
$(go env GOPATH)/bin/file-manager help
```

### Option 3: Using Go Install (Remote Installation)

Once published to a Git repository:

```bash
go install github.com/spacego/file-manager/cmd/filemanager@latest
```

## Usage

```bash
# Show help
./file-manager help

# List files in a directory
./file-manager list /path/to/directory

# Copy a file
./file-manager copy source.txt destination.txt
```

### Examples

```bash
# List current directory contents
./file-manager list .

# Copy a file to backup
./file-manager copy important.txt important_backup.txt

# List root directory
./file-manager list /
```

## Testing

Run all tests:

```bash
go test ./... -v
```

Run tests with coverage:

```bash
go test ./... -cover
```

Run specific package tests:

```bash
go test ./internal/fileops -v
go test ./pkg/manager -v
```

## Development

### Adding New Commands

1. Add the operation to `internal/fileops/ops.go`
2. Wrap it in `pkg/manager/manager.go` if it should be public
3. Add the command handler in `cmd/filemanager/main.go`
4. Write tests for the new functionality

### Code Quality

- Follow Go idioms and best practices
- Use meaningful error messages with context
- Write comprehensive tests for all public functions
- Keep functions small and focused

## License

MIT License

## Author

SpaceGo - Learning Go Programming Journey
