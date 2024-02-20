package cmd

import "testing"

func TestApplyRotationOnLowercase_LeftRotation_Shift3(t *testing.T) {
	char := 'h'
	shift := 3
	rotation := leftRotation
	want := 'e'
	output := applyRotationOnLowercase(char, shift, rotation)
	if want != output {
		t.Fatalf(`applyRotationOnLowercase(%q, %d, %q) = %q, want match for %q`, char, shift, rotation, output, want)
	}
}

func TestApplyRotationOnLowercase_RightRotation_Shift3(t *testing.T) {
	char := 'h'
	shift := 3
	rotation := rightRotation
	want := 'k'
	output := applyRotationOnLowercase(char, shift, rotation)
	if want != output {
		t.Fatalf(`applyRotationOnLowercase(%q, %d, %q) = %q, want match for %q`, char, shift, rotation, output, want)
	}
}

func TestApplyRotationOnUppercase_LeftRotation_Shift3(t *testing.T) {
	char := 'H'
	shift := 3
	rotation := leftRotation
	want := 'E'
	output := applyRotationOnUppercase(char, shift, rotation)
	if want != output {
		t.Fatalf(`applyRotationOnLowercase(%q, %d, %q) = %q, want match for %q`, char, shift, rotation, output, want)
	}
}

func TestApplyRotationOnUppercase_RightRotation_Shift3(t *testing.T) {
	char := 'H'
	shift := 3
	rotation := rightRotation
	want := 'K'
	output := applyRotationOnUppercase(char, shift, rotation)
	if want != output {
		t.Fatalf(`applyRotationOnLowercase(%q, %d, %q) = %q, want match for %q`, char, shift, rotation, output, want)
	}
}

func TestCaesarCipher_RightRotation_Shift23(t *testing.T) {
	text := "ThE QUICK BroWN FOX JUMPS OVER THE LAZY DOG!"
	shift := 23
	rotation := rightRotation
	want := "QeB NRFZH YolTK CLU GRJMP LSBO QEB IXWV ALD!"
	output := applyCaesarCipher(text, shift, rotation)
	if want != output {
		t.Fatalf(`applyCaesarCipher(%q, %d, %q) = %q, want match for %q`, text, shift, rotation, output, want)
	}
}

func TestCaesarCipher_LeftRotation_Shift23(t *testing.T) {
	text := "QeB NRFZH YolTK CLU GRJMP LSBO QEB IXWV ALD!"
	shift := 23
	rotation := leftRotation
	want := "ThE QUICK BroWN FOX JUMPS OVER THE LAZY DOG!"
	output := applyCaesarCipher(text, shift, rotation)
	if want != output {
		t.Fatalf(`applyCaesarCipher(%q, %d, %q) = %q, want match for %q`, text, shift, rotation, output, want)
	}
}
