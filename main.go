package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

// Check if Docker is installed by running "docker --version"
func isDockerInstalled() bool {
	_, err := exec.LookPath("docker")
	return err == nil
}

// Install Docker based on the operating system
func installDocker() error {
	switch runtime.GOOS {
	case "linux":
		// Run Linux Docker installation commands
		cmd := exec.Command("sh", "-c", "curl -fsSL https://get.docker.com -o get-docker.sh && sh get-docker.sh")
		cmd.Stdout = log.Writer()
		cmd.Stderr = log.Writer()
		return cmd.Run()

	case "darwin":
		// For MacOS, use Homebrew to install Docker
		cmd := exec.Command("sh", "-c", "brew install --cask docker")
		cmd.Stdout = log.Writer()
		cmd.Stderr = log.Writer()
		return cmd.Run()

	case "windows":
		// For Windows, prompt the user to install Docker manually or execute a command to install
		fmt.Println("Please install Docker for Windows from https://desktop.docker.com/win/stable/Docker%20Desktop%20Installer.exe")
		// Alternatively, you can use PowerShell commands to install Docker Desktop
		// cmd := exec.Command("powershell", "Start-Process", "DockerDesktopInstaller.exe", "-Wait")
		// cmd.Stdout = log.Writer()
		// cmd.Stderr = log.Writer()
		// return cmd.Run()
		return nil

	default:
		return fmt.Errorf("unsupported OS")
	}
}

func main() {
	// Check if Docker is installed
	if isDockerInstalled() {
		fmt.Println("Docker is already installed")
	} else {
		fmt.Println("Docker is not installed, installing...")
		err := installDocker()
		if err != nil {
			log.Fatalf("Failed to install Docker: %v", err)
		}
		fmt.Println("Docker installed successfully")
	}
}
