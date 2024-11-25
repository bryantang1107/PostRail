package util

import (
	"fmt"
)

func PrintMoves(w int, t string, n1 string, p1 []string, n2 string, p2 []string) {
	fmt.Printf("W=%d, T=%s, N1=%s, P1=%v, N2=%s, P2=%v", w, t, n1, p1, n2, p2)
	fmt.Println("")
}
