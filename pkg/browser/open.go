package browser

import (
	"fmt"
	"os/exec"
	"runtime"
)

// Open opens the specified URL in the default browser.
func Open(url string) error {
	// Determine which command to use based on the platform
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		// First try xdg-open (standard for most Linux distributions)
		if _, err := exec.LookPath("xdg-open"); err == nil {
			cmd = exec.Command("xdg-open", url)
		} else if _, err := exec.LookPath("wslview"); err == nil {
			// For Windows Subsystem for Linux
			cmd = exec.Command("wslview", url)
		} else {
			return fmt.Errorf("no suitable browser command found on Linux")
		}
	case "darwin":
		// macOS
		cmd = exec.Command("open", url)
	case "windows":
		// Windows
		cmd = exec.Command("cmd", "/c", "start", url)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	return cmd.Run()
}