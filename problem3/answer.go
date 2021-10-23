package main

import (
	"flag"
	"fmt"
)

func main() {
	cakePtr := flag.Int("cake", 20, "Cake value")
	applePtr := flag.Int("apple", 25, "Apple value")
	flag.Parse()

	count := BoxesCount(*cakePtr, *applePtr)
	fmt.Println("boxes Ainun can make: ", count)

	cake, aple := InEachBox(*cakePtr, *applePtr)
	fmt.Printf("there are %d cake and %d apple in each box", cake, aple)
}

func BoxesCount(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func InEachBox(cake, aple int) (int, int) {
	divider := BoxesCount(cake, aple)

	return cake / divider, aple / divider
}
