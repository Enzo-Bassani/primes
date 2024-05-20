package fermat

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"primes/globals"
	"primes/linearcongruentialgenerator"
	"testing"
)

// Generate 10⁶ random numbers from 1 to 2⁶³ and compare the developed
// Fermat test results with results from golang official math library.
func TestIsPrimeRandomly(t *testing.T) {
	max := big.NewInt(1 << 62)
	var tests int = 1e5
	for i := 0; i < tests; i++ {
		n, _ := rand.Int(rand.Reader, max)
		got, _ := ProbablyPrime(n, 20)
		want := n.ProbablyPrime(20)
		if got != want {
			t.Errorf("IsPrime() = %v, want %v", got, want)
		}
	}
}

func TestIsPrime(t *testing.T) {
	n := big.NewInt(221)
	
	got, _ := ProbablyPrime(n, 2)
	want := false

	if got != want {
		t.Errorf("IsPrime() = %v, want %v", got, want)
	}
}

func TestIsPrimeCarmichael(t *testing.T) {
	var n big.Int
	n.SetString("561", 10)
	got, _ := ProbablyPrime(&n, 20)
	want := false

	if got != want {
		t.Errorf("ProbablyPrime() = %v, want %v", got, want)
	}

	n.SetString("1105", 10)
	got, _ = ProbablyPrime(&n, 20)
	want = false

	if got != want {
		t.Errorf("ProbablyPrime() = %v, want %v", got, want)
	}
}
func BenchmarkIsPrime(b *testing.B) {
	var seed uint64 = 16086232361280787691
	var m uint64 = 1 << 48 // 2⁴⁸
	var a uint64 = 25214903917
	var c uint64 = 11

	var totalIterations, testedNumbers, oddNumbersTested uint64 = 0, 0, 0

	lcg := linearcongruentialgenerator.NewLCG(a, c, m, seed)
	file, _ := os.Create("primesFermat.txt")
	defer file.Close()

	for _, bits := range globals.TestSizes {
		b.Run(fmt.Sprintf("bits %d", bits), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for {
					n := lcg.Next(bits, false)
					_isPrime, iterations := ProbablyPrime(n, 20)
					totalIterations += uint64(iterations)
					if iterations != 0 {
						oddNumbersTested++
					}
					testedNumbers++
					if _isPrime {
						file.WriteString(n.String() + "\n")
						break
					}
				}
            }
        })
	}

	fmt.Println(fmt.Sprintln("Total iterations: ", totalIterations))
	fmt.Println(fmt.Sprintln("tested numbers: ", testedNumbers))
	fmt.Println(fmt.Sprintln("oddNumbersTested: ", oddNumbersTested))
	fmt.Println(fmt.Sprintln("Average iterations: ", totalIterations / oddNumbersTested))

	file.WriteString(fmt.Sprintln("Total iterations: ", totalIterations))
	file.WriteString(fmt.Sprintln("tested numbers: ", testedNumbers))
	file.WriteString(fmt.Sprintln("oddNumbersTested: ", oddNumbersTested))
	file.WriteString(fmt.Sprintln("Average iterations: ", totalIterations / oddNumbersTested))
}

