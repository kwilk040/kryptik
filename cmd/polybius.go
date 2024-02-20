package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
)

var grid = [][]rune{
	{'A', 'B', 'C', 'D', 'E'},
	{'F', 'G', 'H', 'I', 'K'},
	{'L', 'M', 'N', 'O', 'P'},
	{'Q', 'R', 'S', 'T', 'U'},
	{'V', 'W', 'X', 'Y', 'Z'},
}

var polybiusCmd = &cobra.Command{
	Use:   "polybius",
	Short: "Applies Polybius cipher encrypt/decrypt on message",
	Long: `Example usage:

./kryptik polybius -m "23 15 31 31 34 ,  52 34 42 31 14 \!" -d`,
	Run: func(cmd *cobra.Command, args []string) {
		message, err := cmd.Flags().GetString("message")
		if err != nil {
			e := errors.New(fmt.Sprintf(`could not extract value from "message" flag: %s`, err))
			log.Fatal(e)
		}

		decrypt, err := cmd.Flags().GetBool("decrypt")
		if err != nil {
			e := errors.New(fmt.Sprintf(`could not extract value from "decrypt" flag: %s`, err))
			log.Fatal(e)
		}

		fmt.Printf("Original message: [%s].\n", message)

		if decrypt {
			decryptedMessage := polybiusDecrypt(message)
			fmt.Printf("Decrypted message: [%s].\n", decryptedMessage)
		} else {
			encryptedMessage := polybiusEncrypt(message)
			fmt.Printf("Encrypted message: [%s].\n", encryptedMessage)
		}
	},
}

func init() {
	rootCmd.AddCommand(polybiusCmd)
	polybiusCmd.Flags().StringP("message", "m", "", "Message to encrypt/decrypt")
	polybiusCmd.Flags().BoolP("decrypt", "d", false, "If set, message will be decrypted")
	polybiusCmd.MarkFlagRequired("message")
}

func polybiusEncrypt(message string) string {
	output := strings.Builder{}
	message = strings.TrimSpace(message)
	uppercaseMessage := strings.ToUpper(message)

	for _, char := range uppercaseMessage {
		if char == 'J' {
			char = 'I'
		}

		if char < 'A' || char > 'Z' {
			if char == 32 {
				output.WriteRune(char)
				continue
			}
			output.WriteString(fmt.Sprintf("%c ", char))
			continue
		}

		row, col := findLetterPosition(char)
		if row == -1 || col == -1 {
			output.WriteRune(char)
		} else {
			output.WriteString(fmt.Sprintf("%d%d ", row+1, col+1))
		}
	}

	return strings.TrimSpace(output.String())
}

func polybiusDecrypt(message string) string {
	output := strings.Builder{}
	message = strings.TrimSpace(message)

	split := strings.Split(message, " ")
	for _, v := range split {
		if len(v) == 2 {
			rowCol := strings.Split(v, "")
			row, err := strconv.Atoi(rowCol[0])
			if err != nil {
				e := errors.New(fmt.Sprintf(`could not convert [%s] to integer: [%s]`, rowCol[0], err))
				log.Fatal(e)
			}

			col, err := strconv.Atoi(rowCol[1])
			if err != nil {
				e := errors.New(fmt.Sprintf(`could not convert [%s] to integer: [%s]`, rowCol[1], err))
				log.Fatal(e)
			}

			row = row - 1
			col = col - 1
			if row == 1 && col == 3 {
				output.WriteString("[I/J]")
				continue
			}

			output.WriteRune(grid[row][col])
		} else if v == "" {
			output.WriteString(" ")
		} else {
			output.WriteString(v)
		}
	}
	return output.String()
}

func findLetterPosition(char rune) (int, int) {
	for i, row := range grid {
		for j, val := range row {
			if val == char {
				return i, j
			}
		}
	}
	return -1, -1
}
