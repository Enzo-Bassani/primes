package blumblumshub

import (
	"fmt"
	"math/big"
	"primes/globals"
	"testing"
)



func TestBBS_Next(t *testing.T) {
	p := big.NewInt(7)
	q := big.NewInt(19)
	seed := big.NewInt(100)

	var expectedResult big.Int
	expectedResult.SetString("10010110", 2)

	bbs := NewBBS(seed, p, q)
	output := bbs.Next(8)

	if output.Cmp(&expectedResult) != 0 {
		t.Errorf("BBS.Next() = %v, want %v", output.Text(2), expectedResult.Text(2))
	}
}

func TestBBS_Next2(t *testing.T) {
	p := big.NewInt(30000000091)
	q := big.NewInt(40000000003)
	seed := big.NewInt(5956326528)

	var expectedResult big.Int
	expectedResult.SetString("00011000", 2)

	bbs := NewBBS(seed, p, q)
	output := bbs.Next(8)

	if output.Cmp(&expectedResult) != 0 {
		t.Errorf("BBS.Next() = %v, want %v", output.Text(2), expectedResult.Text(2))
	}
}

func BenchmarkBBS_Next(b *testing.B) {
	p := big.NewInt(30000000091)
	q := big.NewInt(40000000003)
	seed := big.NewInt(5956326528)
	bbs := NewBBS(seed, p, q)

	for _, bits := range globals.TestSizes {	
		b.Run(fmt.Sprintf("bits %d", bits), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                bbs.Next(bits)
            }
        })
	}
}

func BenchmarkBBS_Next_Small(b *testing.B) {
	p := big.NewInt(7)
	q := big.NewInt(19)
	seed := big.NewInt(100)
	
	bbs := NewBBS(seed, p, q)

	for _, bits := range globals.TestSizes {	
		b.Run(fmt.Sprintf("bits %d", bits), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                bbs.Next(bits)
            }
        })
	}
}
