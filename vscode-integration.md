# VSCode Integration Guide for Blamepr

This guide explains how to integrate the `blamepr` tool with Visual Studio Code using keybindings and tasks.

## Prerequisites

1. Install the `blamepr` tool as described in the README.md
2. Ensure your GitHub authentication is properly set up (GITHUB_TOKEN or .netrc)

## Setup

### 1. Configure Keyboard Shortcuts

Add these keybindings to your `keybindings.json` file (File > Preferences > Keyboard Shortcuts > Open Keyboard Shortcuts (JSON)):

```json
[
  {
    "key": "ctrl+shift+b",
    "command": "workbench.action.terminal.sendSequence",
    "args": {
      "text": "blamepr ${file}:${lineNumber}\n"
    },
    "when": "editorTextFocus"
  },
  {
    "key": "ctrl+shift+o",
    "command": "workbench.action.terminal.sendSequence",
    "args": {
      "text": "blamepr -open ${file}:${lineNumber}\n"
    },
    "when": "editorTextFocus"
  }
]
```

### 2. Add Tasks

Add these tasks to your `.vscode/tasks.json` file:

```json
{
  "version": "2.0.0",
  "tasks": [
    {
      "label": "Blame PR for current file",
      "type": "shell",
      "command": "blamepr ${file}",
      "presentation": {
        "reveal": "always",
        "panel": "new"
      },
      "problemMatcher": []
    },
    {
      "label": "Blame PR for current line",
      "type": "shell",
      "command": "blamepr ${file}:${lineNumber}",
      "presentation": {
        "reveal": "always",
        "panel": "new"
      },
      "problemMatcher": []
    },
    {
      "label": "Open PR for current line in browser",
      "type": "shell",
      "command": "blamepr -open ${file}:${lineNumber}",
      "presentation": {
        "reveal": "always",
        "panel": "new"
      },
      "problemMatcher": []
    },
    {
      "label": "View PR details with gh CLI",
      "type": "shell",
      "command": "blamepr -id ${file}:${lineNumber} | xargs gh pr view",
      "presentation": {
        "reveal": "always",
        "panel": "new"
      },
      "problemMatcher": []
    }
  ]
}
```

### 3. Optional: Create a Status Bar Command

For a one-click solution, you can install the "Command Runner" extension from the VSCode marketplace and add this configuration to your `settings.json`:

```json
"command-runner.commands": {
  "Blame PR for current line": "blamepr ${file}:${lineNumber}",
  "Open PR for current line": "blamepr -open ${file}:${lineNumber}"
},
```

## Usage

- Use keyboard shortcuts:
  - `Ctrl+Shift+B` to show PR information for the current line
  - `Ctrl+Shift+O` to open the PR in your browser
  
- Or run tasks:
  - Press `Ctrl+Shift+P`, type "Tasks: Run Task", and select one of the blamepr tasks
  
- Or use the status bar commands (if configured)

## Troubleshooting

- If commands fail, ensure `blamepr` is properly installed and in your PATH
- Check that your terminal is correctly configured in VSCode
- Verify GitHub authentication is set up correctly