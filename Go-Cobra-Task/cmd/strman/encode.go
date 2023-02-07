package strman

import (
	"fmt"
	strman "strman/pkg"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var encodeFormat string
var allowedEncodings = []string{"binary", "ascii", "hex", "base64"}

var encodeCmd = &cobra.Command{
	Use:     "encode",
	Aliases: []string{"en"},
	Short:   "Encodes the string into specific format",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if slices.Contains(allowedEncodings, encodeFormat) {
			res := strman.EncodeString(args[0], encodeFormat)
			fmt.Println(res)
			return nil
		}
		return fmt.Errorf("wrong/unknown format specified")
	},
}

func init() {
	encodeCmd.Flags().StringVarP(&encodeFormat, "format", "f", "", "The format to be used for encoding")
	rootCmd.AddCommand(encodeCmd)
}
