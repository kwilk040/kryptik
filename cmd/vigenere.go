package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var vigenereCmd = &cobra.Command{
	Use:   "vigenere",
	Short: "Applies Vigenere cipher on message",
	Long: `Example usage:

./kryptik vigenere -m "attackatdawn" -k "lemon"`,
	Run: func(cmd *cobra.Command, args []string) {
		message, err := cmd.Flags().GetString("message")
		if err != nil {
			e := errors.New(fmt.Sprintf(`could not extract value from "message" flag: %s`, err))
			log.Fatal(e)
		}

		key, err := cmd.Flags().GetString("key")
		if err != nil {
			e := errors.New(fmt.Sprintf(`could not extract value from "key" flag: %s`, err))
			log.Fatal(e)
		}

		decrypt, err := cmd.Flags().GetBool("decrypt")
		if err != nil {
			e := errors.New(fmt.Sprintf(`could not extract value from "key" flag: %s`, err))
			log.Fatal(e)
		}

		fmt.Printf("Original message: [%s].\nKey: [%s]\n", message, key)

		if decrypt {
			decryptedMessage := vigenereDecrypt(message, key)
			fmt.Printf("Decrypted message: [%s].\n", decryptedMessage)
		} else {
			encryptedMessage := vigenereEncrypt(message, key)
			fmt.Printf("Encrypted message: [%s].\n", encryptedMessage)
		}
	},
}

func init() {
	rootCmd.AddCommand(vigenereCmd)
	vigenereCmd.Flags().StringP("message", "m", "", "Message to encrypt/decrypt")
	vigenereCmd.Flags().StringP("key", "k", "", "Key used for encryption/decryption")
	vigenereCmd.Flags().BoolP("decrypt", "d", false, "If set, message will be decrypted")
	vigenereCmd.MarkFlagRequired("message")
	vigenereCmd.MarkFlagRequired("key")
	vigenereCmd.MarkFlagsRequiredTogether("message", "key")
}

func vigenereEncrypt(message, key string) string {
	message = strings.TrimSpace(strings.ToUpper(message))
	key = strings.TrimSpace(strings.ToUpper(key))
	output := strings.Builder{}

	for i, char := range message {
		if char >= 'A' && char <= 'Z' {
			shifted := (char-'A'+int32(key[i%len(key)]-'A'))%26 + 'A'
			output.WriteRune(shifted)
		} else {
			output.WriteRune(char)
		}
	}
	return output.String()
}

func vigenereDecrypt(message, key string) string {
	message = strings.TrimSpace(strings.ToUpper(message))
	key = strings.TrimSpace(strings.ToUpper(key))
	output := strings.Builder{}

	for i, char := range message {
		if char >= 'A' && char <= 'Z' {
			shifted := (char-'A'+26-(int32(key[i%len(key)]-'A')))%26 + 'A'
			output.WriteRune(shifted)
		} else {
			output.WriteRune(char)
		}
	}
	return output.String()
}
