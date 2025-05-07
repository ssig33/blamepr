## **blamepr: A GitHub PR Trace CLI Tool**

### **Overview**

`blamepr` is a command-line tool that maps source code lines or files to the GitHub Pull Request (PR) that last modified them. It enables quick traceability between local code and the associated change history on GitHub.

---

### **Use Cases**

* `blamepr path/to/file.go`
  → Show the latest PR ID that modified this file.

* `blamepr path/to/file.go:123`
  → Show the PR ID that last modified line 123.

* `blamepr -open path/to/file.go:123`
  → Open the PR in the default browser.

---

### **Key Features**

* Resolves local file and line to Git commit.
* Maps Git commit to PR via GitHub API.
* Supports opening PR in browser.
* Works with GitHub-hosted repos (via HTTPS or SSH).

---

### **Implementation Outline**

* **Language**: Go
* **Inputs**: File path (with optional line number), flags
* **Logic**:

  1. Parse file path and optional line number.
  2. Use `git blame` to find the commit hash.
  3. Query GitHub API to find the PR associated with that commit.
  4. Output the PR number, title, and URL.
  5. If `-open` flag is set, open the PR URL in a browser.

---

### **Dependencies**

* Git CLI (`git blame`, `git rev-parse`, etc.)
* GitHub API token via `GITHUB_TOKEN` env var or .netrc file.
* `xdg-open`, `wsl-view`, or `open` for browser integration (platform-dependent)

