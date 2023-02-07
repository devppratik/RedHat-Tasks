package strman

import (
	"fmt"
	strman "strman/pkg"

	"github.com/spf13/cobra"
)

var separator string
var concatCmd = &cobra.Command{
	Use:     "concat",
	Aliases: []string{"ct"},
	Short:   "Concats the given inputs",
	Long:    "Concats one or more of the given inputs provided and displays it on the output. Use the separator flag to define the separator between strings. Defaults to none",
	Args:    cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		res := strman.ConcatStrings(args, separator)
		fmt.Println(res)
	},
}

func init() {
	concatCmd.Flags().StringVarP(&separator, "sep", "s", "", "Defines the separator between concatenated strings. Defaults to none")
	rootCmd.AddCommand(concatCmd)
}
