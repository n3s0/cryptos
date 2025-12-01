package cmd

import (
	"os"
	"fmt"
	
	"github.com/spf13/cobra"
)

var verbose bool

var rootCmd = &cobra.Command{
	Use: "cryptos",
	Short: "CLI application that will decode/decrypt multiple ciphers and encryption alorithms",
	Long: "CLI application that will decode/decrypt multiple ciphers and encryption alorithms",
	Run: func (cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", "v", false, "Make output more verbose")
}
