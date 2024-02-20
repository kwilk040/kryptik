package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

type rotation string

func (e *rotation) Type() string {
	return "rotation"
}

func (e *rotation) String() string {
	return string(*e)
}

func (e *rotation) Set(v string) error {
	switch v {
	case "left", "right":
		*e = rotation(v)
		return nil
	default:
		return errors.New(`must be one of "left" or "right".'`)
	}
}

const (
	leftRotation  rotation = "left"
	rightRotation rotation = "right"
)

var (
	Message  string
	Shift    int
	Rotation rotation
)

var caesarCmd = &cobra.Command{
	Use:   "caesar",
	Short: "Applies Caesar cipher on message",
	Long: `Example usage:

./kryptik caesar -m "Hello" -s 3 -r right`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Original message: [%s].\n", Message)
		var caesarMessage = applyCaesarCipher(Message, Shift, Rotation)
		fmt.Printf("Message after applying [%s] rotation with shift [%d]: [%s].\n", Rotation, Shift, caesarMessage)
	},
}

func init() {
	rootCmd.AddCommand(caesarCmd)
	caesarCmd.Flags().StringVarP(&Message, "message", "m", "", "Message to encrypt/decrypt")
	caesarCmd.Flags().IntVarP(&Shift, "shift", "s", 0, "Value of shift")
	caesarCmd.Flags().VarP(&Rotation, "rotation", "r", `Sets shift direction. Allowed values: "right", "left"`)
	caesarCmd.MarkFlagRequired("message")
	caesarCmd.MarkFlagRequired("shift")
	caesarCmd.MarkFlagRequired("rotation")
	caesarCmd.MarkFlagsRequiredTogether("message", "shift", "rotation")
}

func applyCaesarCipher(message string, shift int, rot rotation) string {
	output := strings.Builder{}
	message = strings.TrimSpace(message)
	for _, char := range message {
		if char >= 'a' && char <= 'z' {
			output.WriteRune(applyRotationOnLowercase(char, shift, rot))
			continue
		}
		if char >= 'A' && char <= 'Z' {
			output.WriteRune(applyRotationOnUppercase(char, shift, rot))
			continue
		}
		output.WriteRune(char)
	}
	return output.String()
}

func applyRotationOnLowercase(char rune, shift int, rot rotation) rune {
	switch rot {
	case leftRotation:
		return (char-'a'-rune(shift)+26)%26 + 'a'

	case rightRotation:
		return (char-'a'+rune(shift)+26)%26 + 'a'
	default:
		return rune(0)
	}
}

func applyRotationOnUppercase(char rune, shift int, rot rotation) rune {
	switch rot {
	case leftRotation:
		return (char-'A'-rune(shift)+26)%26 + 'A'

	case rightRotation:
		return (char-'A'+rune(shift)+26)%26 + 'A'
	default:
		return rune(0)
	}
}
