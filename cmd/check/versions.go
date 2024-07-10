/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package check

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// versionsCmd represents the versions command
var versionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "Verifique as versões das ferramentas instaladas",
	Long: `Este comando verifica e imprime as versões do Terraform, Docker, AWS CLI, Kubernetes (kubectl), OpenTofu e VSCode
se eles estiverem instalados na sua máquina.`,
	Run: func(cmd *cobra.Command, args []string) {
		osType := runtime.GOOS
		osDistro := getLinuxDistro()
		fmt.Printf("Sistema Operacional detectado: %s\n", osType)
		if osType == "linux" {
			fmt.Printf("Distribuição Linux: %s\n", osDistro)
		}

		var installed []string
		var notFound []string

		tools := []string{"terraform", "docker", "aws", "kubectl", "opentofu", "vscode"}
		for _, tool := range tools {
			version, err := checkVersion(tool, osType)
			if err != nil {
				notFound = append(notFound, fmt.Sprintf("\033[1m%s\033[0m não está instalado ou não foi encontrado no PATH", strings.Title(tool)))
			} else {
				installed = append(installed, fmt.Sprintf("\033[1m%s\033[0m versão: %s", strings.Title(tool), version))
			}
		}

		fmt.Println("\nFerramentas Instaladas:")
		for _, msg := range installed {
			fmt.Println(msg)
		}

		fmt.Println("\nFerramentas Não Encontradas:")
		for _, msg := range notFound {
			fmt.Println(msg)
		}
	},
}

func init() {
	CheckCmd.AddCommand(versionsCmd)
}

// checkVersion runs the provided command to check its version
func checkVersion(tool, osType string) (string, error) {
	var versionCmd *exec.Cmd

	switch tool {
	case "terraform":
		versionCmd = exec.Command("terraform", "version")
	case "docker":
		if osType == "windows" {
			versionCmd = exec.Command("docker", "version", "--format", "{{.Server.Version}}")
		} else {
			versionCmd = exec.Command("docker", "--version")
		}
	case "aws":
		versionCmd = exec.Command("aws", "--version")
	case "kubectl":
		versionCmd = exec.Command("kubectl", "version", "--client", "--short")
	case "opentofu":
		versionCmd = exec.Command("opentofu", "version")
	case "vscode":
		versionCmd = exec.Command("code", "--version")
	default:
		return "", fmt.Errorf("ferramenta desconhecida: %s", tool)
	}

	output, err := versionCmd.Output()
	if err != nil {
		return "", err
	}

	version := strings.TrimSpace(string(output))
	return version, nil
}

// getLinuxDistro detects and returns the Linux distribution
func getLinuxDistro() string {
	output, err := exec.Command("cat", "/etc/os-release").Output()
	if err != nil {
		return "Distribuição não identificada"
	}

	var distro string
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			distro = strings.TrimPrefix(line, "PRETTY_NAME=")
			distro = strings.Trim(distro, "\"")
			break
		}
	}

	if distro == "" {
		return "Distribuição não identificada"
	}

	return distro
}

