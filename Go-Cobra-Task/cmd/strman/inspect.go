package strman

import (
	"fmt"
	strman "strman/pkg"

	"github.com/spf13/cobra"
)

var onlyDigits bool
var inspectCmd = &cobra.Command{
	Use:     "inspect",
	Aliases: []string{"insp"},
	Short:   "Inspects a string",
	Long:    "Inspects a string & returns its length. Passing the digit flag analyzes the number of digits in the string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		i := args[0]
		res, kind := strman.Inspect(i, onlyDigits)

		pluralS := "s"
		if res == 1 {
			pluralS = ""
		}
		fmt.Printf("'%s' has %d %s%s.\n", i, res, kind, pluralS)
	},
}

func init() {
	inspectCmd.Flags().BoolVarP(&onlyDigits, "digits", "d", false, "Count only digits")
	rootCmd.AddCommand(inspectCmd)
}
