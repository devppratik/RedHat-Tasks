package strman

import (
	"fmt"
	strman "strman/pkg"

	"github.com/spf13/cobra"
)

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Generates a random string of given length",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := strman.RandomString(args[0])
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}
