package cmd

import "testing"

func TestTrithemiusEncryption(t *testing.T) {
	text := "HELLO"
	want := "HFNOS"
	output := trithemiusEncrypt(text)
	if want != output {
		t.Fatalf(`trithemiusEncrypt(%q) = %q, want match for %q`, text, output, want)
	}
}

func TestTrithemiusDecryption(t *testing.T) {
	text := "HFNOS"
	want := "HELLO"
	output := trithemiusDecrypt(text)
	if want != output {
		t.Fatalf(`trithemiusDecrypt(%q) = %q, want match for %q`, text, output, want)
	}
}
