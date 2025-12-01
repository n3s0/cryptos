package cmd

import (
	"fmt"
	"os"
	
	"github.com/spf13/cobra"
	"github.com/n3s0/cryptos/utils"
)

var decodeHex bool
var decodeBase64 bool
var decodeRot13 bool
var inputString string

var decodeCmd = &cobra.Command{
	Use: "decode",
	Short: "Sub command used to decode various strings",
	Long: "Sub command used to decode cipher text or strings like base64 to hex.",
	Run: func (cmd *cobra.Command, args []string) {
		if inputString == "" {
			fmt.Errorf("[!] Needs input. Use --input or -i to send input! Exiting!\n")
			os.Exit(1)
		}
		
		if decodeHex {
			fmt.Printf("[i] Decoding hex to string...\n")
			decodedHexString, err := utils.DecodeHexToString(inputString)
			if err != nil {
				fmt.Errorf("[!] Error: %v\n", err)
			}

			fmt.Printf("[+] Encoded: %s\n", inputString)
			fmt.Printf("[+] Decoded: %s\n", decodedHexString)
		}

		if decodeBase64 {
			fmt.Printf("[i] This flag will decode Base64 to string...\n")
			decodedBase64String, err := utils.DecodeBase64ToString(inputString)
			if err != nil {
				fmt.Errorf("[!] Error: %v\n", err)
			}

			fmt.Printf("[+] Encoded: %s\n", inputString)
			fmt.Printf("[+] Decoded: %s\n", decodedBase64String)
		}

		if decodeRot13 {
			decodedRot13String, err := utils.DecodeRot13(inputString)
			if err != nil {
				fmt.Errorf("[!] Error: %v\n", err)
			}
			
			fmt.Printf("[+] Just some output: %s", decodedRot13String)
		}
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)
	decodeCmd.Flags().BoolVarP(&decodeHex, "hex", "", false, "Decodes Hex to string")
	decodeCmd.Flags().BoolVarP(&decodeBase64, "base64", "", false, "Decodes Base64 to string")
	decodeCmd.Flags().BoolVarP(&decodeRot13, "rot13", "", false, "Decodes ROT13")
	decodeCmd.Flags().StringVarP(&inputString, "input", "i", "", "Accepts string from user")
}

