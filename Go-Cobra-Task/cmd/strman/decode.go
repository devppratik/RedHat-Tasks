package strman

import (
	"fmt"
	strman "strman/pkg"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var decodeFormat string
var decodeCmd = &cobra.Command{
	Use:     "decode",
	Aliases: []string{"en"},
	Short:   "Decodes the string into UTF-8 from the specific format",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if slices.Contains(allowedEncodings, decodeFormat) {
			res := strman.DecodeString(args[0], decodeFormat)
			fmt.Println(res)
		}
		return fmt.Errorf("wrong/unknown format specified")

	},
}

func init() {
	decodeCmd.Flags().StringVarP(&decodeFormat, "format", "f", "", "The format to be used for decoding")
	rootCmd.AddCommand(decodeCmd)
}
