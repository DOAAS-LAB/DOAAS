/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package update

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/spf13/cobra"
)

var toolsToUpdate = []string{"tofu", "docker", "aws", "kubectl", "opentofu", "vscode"}

// updateCmd represents the update command
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update tools installed",
	Long:  `Check for available updates and update selected tools.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Verifying tools with pending updates...")
		fmt.Println()

		// Simulate checking for updates
		pendingUpdates := checkPendingUpdates()

		if len(pendingUpdates) == 0 {
			fmt.Println("Nenhum update disponível.")
			return
		}

		// Display tools with pending updates
		fmt.Println("Ferramentas com updates pendentes:")
		for i, tool := range pendingUpdates {
			fmt.Printf("%d. %s\n", i+1, tool)
		}
		fmt.Println()

		// Prompt user to select tools for update
		selectedTools := promptForUpdates()

		// Install selected tools
		fmt.Println("\nIniciando instalação dos updates...")

		for _, toolIndex := range selectedTools {
			if toolIndex < 1 || toolIndex > len(pendingUpdates) {
				fmt.Printf("Opção inválida: %d. Ignorando...\n", toolIndex)
				continue
			}

			tool := pendingUpdates[toolIndex-1]
			fmt.Printf("Instalando %s...\n", tool)
			installAndUpdate(tool)
		}

		fmt.Println("\nTodos os updates foram instalados com sucesso!")
	},
}

func init() {
	// Initialize the updateCmd and add it to the root command or another appropriate command.
	// For example:
	// rootCmd.AddCommand(update.UpdateCmd)
}

// checkPendingUpdates simulates checking for pending updates
func checkPendingUpdates() []string {
	// Here you would typically implement logic to check for pending updates for each tool.
	// For demonstration purposes, we return a static list.
	return []string{"tofu", "docker", "aws", "kubectl", "opentofu", "vscode"}
}

// promptForUpdates prompts the user to select tools for update
func promptForUpdates() []int {
	var selectedTools []int

	for {
		var input string
		fmt.Print("Selecione o número da ferramenta para atualizar (0 para finalizar): ")
		fmt.Scanln(&input)

		if input == "0" {
			break
		}

		toolIndex, err := strconv.Atoi(input)
		if err != nil || toolIndex < 0 {
			fmt.Println("Opção inválida. Por favor, selecione um número válido.")
			continue
		}

		selectedTools = append(selectedTools, toolIndex)
	}

	return selectedTools
}

// installAndUpdate installs and updates a tool based on OS type
func installAndUpdate(tool string) {
	osType := runtime.GOOS
	switch osType {
	case "windows":
		installWindows(tool)
	case "linux":
		instalFedora(tool)
	default:
		fmt.Printf("Sistema Operacional não suportado: %s\n", osType)
	}
}

// instalFedora installs and updates a tool on Linux (Fedora) using dnf
func instalFedora(tool string) {
	// Example dnf command to install/update tools on Fedora
	var cmd *exec.Cmd

	switch tool {
	case "tofu":
		cmd = exec.Command("dnf", "install", "tofu", "-y")
	case "docker":
		cmd = exec.Command("dnf", "install", "docker-ce", "-y")
	case "aws":
		cmd = exec.Command("dnf", "install", "awscli", "-y")
	case "kubectl":
		cmd = exec.Command("dnf", "install", "kubectl", "-y")
	case "opentofu":
		// Replace with actual installation command for opentofu on Fedora
		fmt.Printf("Installation command for %s on Fedora\n", tool)
	case "vscode":
		cmd = exec.Command("dnf", "install", "code", "-y")
	default:
		fmt.Printf("Ferramenta desconhecida: %s\n", tool)
		return
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Erro ao instalar %s no Linux (Fedora): %v\n", tool, err)
		return
	}

	fmt.Printf("%s instalado e atualizado com sucesso no Linux (Fedora)!\n", tool)
	fmt.Println(string(output))
}
