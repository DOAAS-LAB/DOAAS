package cmd

import (
    "fmt"
    "os/exec"
    "strings"
    "github.com/spf13/cobra"
)

func getVersion(cmd string, args ...string) (string, error) {
    output, err := exec.Command(cmd, args...).Output()
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(string(output)), nil
}

func checkUpdates(program, currentVersion string) (bool, error) {
    // Implementar lógica para verificar atualizações
    // Esta função deve retornar true se houver uma atualização disponível
    return false, nil
}

var checkCmd = &cobra.Command{
    Use:   "check",
    Short: "Verificar versões do Terraform, AWS CLI e Docker",
    Run: func(cmd *cobra.Command, args []string) {
        programs := map[string]string{
            "terraform": "terraform version",
            "aws":       "aws --version",
            "docker":    "docker --version",
        }

        for name, cmd := range programs {
            version, err := getVersion("sh", "-c", cmd)
            if err != nil {
                fmt.Printf("Erro ao verificar a versão do %s: %v\n", name, err)
                continue
            }

            updateAvailable, err := checkUpdates(name, version)
            if err != nil {
                fmt.Printf("Erro ao verificar atualizações do %s: %v\n", name, err)
                continue
            }

            status := "atualizado"
            if updateAvailable {
                status = "desatualizado"
            }

            fmt.Printf("%s: %s (%s)\n", name, version, status)
        }
    },
}

func init() {
    rootCmd.AddCommand(checkCmd)
}
