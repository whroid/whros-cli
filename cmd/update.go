package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update whros to the latest version",
	RunE:  runUpdate,
}

func runUpdate(cmd *cobra.Command, args []string) error {
	version := "latest"
	if len(args) > 0 {
		version = args[0]
	}

	platform := getPlatform()
	osName := platform[:strings.Index(platform, ":")]
	arch := platform[strings.Index(platform, ":")+1:]
	ext := ""
	if osName == "windows" {
		ext = ".exe"
	}

	repo := "whroid/whros-cli"
	filename := fmt.Sprintf("whros-%s-%s%s", osName, arch, ext)
	downloadURL := fmt.Sprintf("https://github.com/%s/releases/download/%s/%s", repo, version, filename)

	fmt.Printf("Downloading whros %s for %s...\n", version, platform)
	fmt.Printf("URL: %s\n", downloadURL)

	tmpfile, err := os.CreateTemp("", "whros-update-*")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	tmpPath := tmpfile.Name()
	defer os.Remove(tmpPath)
	tmpfile.Close()

	if err := downloadFile(tmpPath, downloadURL); err != nil {
		return fmt.Errorf("failed to download: %w", err)
	}

	if err := os.Chmod(tmpPath, 0755); err != nil {
		return fmt.Errorf("failed to chmod: %w", err)
	}

	execPath, err := os.Executable()
	if err != nil {
		execPath = "/usr/local/bin/whros"
	}

	fmt.Printf("Installing to %s...\n", execPath)

	if err := os.Rename(tmpPath, execPath); err != nil {
		if strings.Contains(err.Error(), "permission") || strings.Contains(err.Error(), "Operation not permitted") {
			fmt.Println("Permission denied. Trying with sudo...")
			sudoTmp := "/tmp/whros-update"
			if err := downloadFile(sudoTmp, downloadURL); err != nil {
				return fmt.Errorf("failed to download: %w", err)
			}
			defer os.Remove(sudoTmp)
			if err := os.Chmod(sudoTmp, 0755); err != nil {
				return fmt.Errorf("failed to chmod: %w", err)
			}
			sudoCmd := exec.Command("sudo", "mv", sudoTmp, execPath)
			if output, err := sudoCmd.CombinedOutput(); err != nil {
				return fmt.Errorf("failed to install with sudo: %w\n%s", err, output)
			}
		} else {
			return fmt.Errorf("failed to replace binary: %w", err)
		}
	}

	fmt.Println("Update complete!")
	fmt.Println("Run 'whros help' to get started.")
	return nil
}

func getPlatform() string {
	os := runtime.GOOS
	arch := runtime.GOARCH

	switch os {
	case "darwin":
		os = "darwin"
	case "linux":
		os = "linux"
	case "windows":
		os = "windows"
	default:
		os = "linux"
	}

	switch arch {
	case "amd64":
		arch = "amd64"
	case "arm64":
		arch = "arm64"
	default:
		arch = "amd64"
	}

	return os + ":" + arch
}

func downloadFile(path, urlStr string) error {
	cmd := exec.Command("curl", "-L", "-o", path, urlStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func init() {
	RootCmd.AddCommand(updateCmd)
}
