package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 48; i < 56; i++ {
		for j := 49; j < 57; j++ {
			for k := 50; k < 58; k++ {
				z01.PrintRune(rune(i))
				z01.PrintRune(rune(j))
				z01.PrintRune(rune(k))
				z01.PrintRune(rune(44))
				z01.PrintRune(rune(32))
			}
		}
	}
	z01.PrintRune('\n')
}
