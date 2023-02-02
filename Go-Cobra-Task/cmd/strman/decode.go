package strman

import (
	"fmt"
	strman "strman/pkg"

	"github.com/spf13/cobra"
)

var decodeFormat string
var decodeCmd = &cobra.Command{
	Use:     "decode",
	Aliases: []string{"en"},
	Short:   "Decodes the string into UTF-8 from the specific format",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := strman.DecodeString(args[0], decodeFormat)
		fmt.Println(res)
	},
}

func init() {
	decodeCmd.Flags().StringVarP(&decodeFormat, "format", "f", "", "The format for hashing")
	rootCmd.AddCommand(decodeCmd)
}
