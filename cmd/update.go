package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
    Use:   "update",
    Short: "Atualizar programas desatualizados",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Função de atualização ainda não implementada")
    },
}

func init() {
    rootCmd.AddCommand(updateCmd)
}
