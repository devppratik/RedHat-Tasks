package strman

import (
	"fmt"
	strman "strman/pkg"

	"github.com/spf13/cobra"
)

var encodeFormat string
var encodeCmd = &cobra.Command{
	Use:     "encode",
	Aliases: []string{"en"},
	Short:   "Encodes the string into specific format",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := strman.EncodeString(args[0], encodeFormat)
		fmt.Println(res)
	},
}

func init() {
	encodeCmd.Flags().StringVarP(&encodeFormat, "format", "f", "", "The format for hashing")
	rootCmd.AddCommand(encodeCmd)
}
