package cmd

import "testing"

func TestVigenereEncryption(t *testing.T) {
	text := "ATTACKING TONIGHT"
	key := "OCULORHINOLARINGOLOGY"
	want := "OVNLQBPVT EOEQTNH"
	output := vigenereEncrypt(text, key)
	if want != output {
		t.Fatalf(`vigenereEncrypt(%q, %q) = %q, want match for %q`, text, key, output, want)
	}
}

func TestVigenereDecryption(t *testing.T) {
	text := "OVNLQBPVT EOEQTNH"
	key := "OCULORHINOLARINGOLOGY"
	want := "ATTACKING TONIGHT"
	output := vigenereDecrypt(text, key)
	if want != output {
		t.Fatalf(`vigenereDecrypt(%q, %q) = %q, want match for %q`, text, key, output, want)
	}
}
