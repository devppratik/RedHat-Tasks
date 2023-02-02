// Hashes string in specified format

package strman

import (
	"fmt"
	strman "strman/pkg"

	"github.com/spf13/cobra"
)

var hashFormat string

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Hashes the string into specific format",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := strman.HashString(args[0], hashFormat)
		fmt.Println(res)
	},
}

func init() {
	hashCmd.Flags().StringVarP(&hashFormat, "format", "f", "", "The format for hashing")
	rootCmd.AddCommand(hashCmd)
}
