package cmd

import (
	"github.com/spf13/cobra"
)


var decryptCmd = &cobra.Command{
	Use: "decode",
	Short: "Sub command used to decode various strings",
	Long: "Sub command used to decode cipher text or strings like base64 to hex.",
	Run: func (cmd *cobra.Command, args []string) {
	},
}
