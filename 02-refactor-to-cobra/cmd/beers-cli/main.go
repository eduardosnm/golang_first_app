// go run cmd/beers-cli/main.go beers -i opcion1

package main

import (
	"github.com/eduardosnm/golang_first_app/02-refactor-to-cobra/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd())
	rootCmd.Execute()
}
