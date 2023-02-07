package strman

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var jsonPrettyCmd = &cobra.Command{
	Use:     "jsonpretty",
	Aliases: []string{"json"},
	Short:   "Pretty Prints the JSON from the file specified",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		src, _ := os.ReadFile(args[0])
		dst := &bytes.Buffer{}
		if err := json.Indent(dst, src, "", "\t"); err != nil {
			panic(err)
		}

		fmt.Println(dst.String())
	},
}

func init() {
	rootCmd.AddCommand(jsonPrettyCmd)
}
