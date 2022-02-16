package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CobraFnShop function definion of run cobra command
type CobraFnShop func(cmd *cobra.Command, args []string)

var shops = map[string]string{
	"01DC9ZAPGKEQJS4P4A48EG3P43": "Mercadona",
	"01DC9ZB23EW0J0ARAER09SJDKC": "Carrefour",
	"01DC9ZB89V1PQD977ZE6QXSQHH": "Alcampo",
}

const idFlagShop = "id"

// InitshopsCmd initialize beers command
func InitShopsCmd() *cobra.Command {
	shopsCmd := &cobra.Command{
		Use:   "shops",
		Short: "Print data about shops",
		Run:   runShopsFn(),
	}

	shopsCmd.Flags().StringP(idFlagShop, "i", "", "id of the shop")

	return shopsCmd
}

func runShopsFn() CobraFnShop {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(idFlagShop)

		if id != "" {
			fmt.Println(shops[id])
		} else {
			fmt.Println(shops)
		}
	}
}
