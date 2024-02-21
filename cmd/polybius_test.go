package cmd

import (
	"testing"
)

func TestPolybiusEncryption(t *testing.T) {
	text := "POLYBIUS SQUARE"
	want := "35 34 31 54 12 24 45 43  43 41 45 11 42 15"
	output := polybiusEncrypt(text)
	if want != output {
		t.Fatalf(`polybiusEncrypt(%q) = %q, want match for %q`, text, output, want)
	}
}

func TestPolybiusDecryption(t *testing.T) {
	text := "35 34 31 54 12 24 45 43  43 41 45 11 42 15"
	want := "POLYB[I/J]US SQUARE"
	output := polybiusDecrypt(text)
	if want != output {
		t.Fatalf(`polybiusDecrypt(%q) = %q, want match for %q`, text, output, want)
	}
}
