/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package strman

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"
var rootCmd = &cobra.Command{
	Use:     "strman",
	Version: version,
	Short:   "strman - a simple CLI to transform and inspect strings",
	Long: `strman is a super fancy CLI     
		One can use strman to modify or inspect strings straight from the terminal`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
