package strman

import (
	"fmt"
	strman "strman/pkg"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var hashFormat string
var allowedFormat = []string{"sha256", "md5", "sha1", "sha512"}

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Hashes the string into specific format",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if slices.Contains(allowedFormat, hashFormat) {
			res := strman.HashString(args[0], hashFormat)
			fmt.Println(res)
			return nil
		}
		return fmt.Errorf("wrong/unknown format specified")
	},
}

func init() {
	hashCmd.Flags().StringVarP(&hashFormat, "format", "f", "", "The format to be used for hashing")
	rootCmd.AddCommand(hashCmd)
}
