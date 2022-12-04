package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/exercises/exercise2-2/weightconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "pk: %v\n", err)
			os.Exit(1)
		}
		lb := weightconv.Pound(t)
		kg := weightconv.Kilogram(t)
		fmt.Printf("%s = %s, %s = %s\n", 
			lb, weightconv.PToK(lb), kg, weightconv.KToP(kg))
	}
}