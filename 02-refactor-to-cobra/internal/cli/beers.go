package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

var beers = map[string]string{
	"opcion1": "Quilmes",
	"opcion2": "Salta",
	"opcion3": "Heineken",
}

const idFlag = "id"

// InitBeersCmd initialize beers command. Return reference to pointer (*cobra.Command)
func InitBeersCmd() *cobra.Command {
	beersCmd := &cobra.Command{ // & create a pointer
		Use:   "beers",
		Short: "Print data about beers",
		Run:   runBeersFn(),
	}

	beersCmd.Flags().StringP(idFlag, "i", "", "id of the beer") // shortcut for id flag

	return beersCmd
}

// Retunr a function
func runBeersFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idFlag)

		if id != "" {
			fmt.Println(beers[id])
		} else {
			fmt.Println(beers)
		}
	}
}
