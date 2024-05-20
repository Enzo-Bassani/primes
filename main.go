package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"primes/blumblumshub"
	"primes/fermat"
	"primes/globals"
	"primes/linearcongruentialgenerator"
	"primes/millerrabin"
)

func main() {
	lcgExample()
	bbsExample()
}

func findDifference() {
	file, _ := os.Create("primesFermat.txt")
	defer file.Close()
	
	var max big.Int
	max.Exp(globals.Two, globals.Bastante, nil)
	for {
		n, _ := rand.Int(rand.Reader, &max)
		miller, _ := millerrabin.ProbablyPrime(n, 8)
		fermat, _ := fermat.ProbablyPrime(n, 8)
		the_truth := n.ProbablyPrime(8)

		if miller != fermat || miller != the_truth {
			file.WriteString(n.String() + " ")
			file.WriteString(fmt.Sprintf("Miller %v", miller))
			file.WriteString(fmt.Sprintf("Fermat %v", fermat))
			file.WriteString(fmt.Sprintf("The_truth %v", the_truth))
		}
	}
}

func lcgExample() {
	var a uint64 = 25214903917
	var c uint64 = 11
	var m uint64 = 1 << 48 // 2⁴⁸
	var seed uint64 = 14583667666143694007
	bits := 1024

	lcg := linearcongruentialgenerator.NewLCG(a, c, m, seed)

	eba := lcg.Next(bits, true)

	fmt.Println(eba)
	fmt.Println(lcg.GeneratedNumbers())
}

func bbsExample() {
	p := blumblumshub.GetBlumPrime()
	q := blumblumshub.GetBlumPrime()
	seed, _ := rand.Prime(rand.Reader, 128)

	fmt.Println("p: ", p)
	fmt.Println("q: ", q)
	fmt.Println("seed: ", seed)

	bbs := blumblumshub.NewBBS(p, q, seed)
	fmt.Println(bbs.Next(1024))
}
