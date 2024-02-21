package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var trithemiusCmd = &cobra.Command{
	Use:   "trithemius",
	Short: "Applies Trithemius cipher on message",
	Long: `Example usage:

./kryptik trithemius -m "HFNOS" -d`,
	Run: func(cmd *cobra.Command, args []string) {
		message, err := cmd.Flags().GetString("message")
		if err != nil {
			e := errors.New(fmt.Sprintf(`could not extract value from "message" flag: %s`, err))
			log.Fatal(e)
		}

		decrypt, err := cmd.Flags().GetBool("decrypt")
		if err != nil {
			e := errors.New(fmt.Sprintf(`could not extract value from "key" flag: %s`, err))
			log.Fatal(e)
		}

		fmt.Printf("Original message: [%s].\n", message)

		if decrypt {
			decryptedMessage := trithemiusDecrypt(message)
			fmt.Printf("Decrypted message: [%s].\n", decryptedMessage)
		} else {
			encryptedMessage := trithemiusEncrypt(message)
			fmt.Printf("Encrypted message: [%s].\n", encryptedMessage)
		}
	},
}

func init() {
	rootCmd.AddCommand(trithemiusCmd)
	trithemiusCmd.Flags().StringP("message", "m", "", "Message to encrypt/decrypt")
	trithemiusCmd.Flags().BoolP("decrypt", "d", false, "If set, message will be decrypted")
	trithemiusCmd.MarkFlagRequired("message")
}

func trithemiusDecrypt(message string) string {
	return vigenereDecrypt(message, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

func trithemiusEncrypt(message string) string {
	return vigenereEncrypt(message, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}
