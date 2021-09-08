package piscine

import "github.com/01-edu/z01"

// Write a function that prints,
// in ascending order and on a single line:
// all unique combinations of three different digits so that,
// the first digit is lower than the second, and the second is lower than the third.
// These combinations are separated by a comma and a space.

func PrintComb() {
	for i := 48; i < 56; i++ {
		for j := 49; j < 57; j++ {
			for k := 51; k < 58; k++ {
				z01.PrintRune(rune(i))
				z01.PrintRune(rune(j))
				z01.PrintRune(rune(k))
				z01.PrintRune(rune(44))
				z01.PrintRune(rune(32))
			}
		}
	}
}
