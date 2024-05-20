package linearcongruentialgenerator

import (
	"fmt"
	"math/big"
	"primes/globals"
	"reflect"
	"testing"
)

// This test follows the example from https://www.geeksforgeeks.org/linear-congruence-method-for-generating-pseudo-random-numbers/
func TestLCG_Next(t *testing.T) {
	var seed uint64 = 5
	var m uint64 = 7
	var a uint64 = 3
	var c uint64 = 3

	numbersToGenerate := 9
	bits := numbersToGenerate * 32

	lcg := NewLCG(a, c, m, seed)
	lcg.Next(bits, true)
	
	expectedGeneratedNumbers := []uint64{4, 1, 6, 0, 3, 5, 4, 1, 6}
	if !reflect.DeepEqual(lcg.GeneratedNumbers(), expectedGeneratedNumbers) {
		t.Errorf("LCG.Next() = %v, want %v", lcg.GeneratedNumbers(), expectedGeneratedNumbers)
	}
}

func TestLCG_Next2(t *testing.T) {
	var a uint64 = 25214903917
	var c uint64 = 11
	var m uint64 = 1 << 48 // 2⁴⁸
	var seed uint64 = 14583667666143694007
	bits := 1024

	var expectedResult big.Int
	expectedResult.SetString("33703635961449644548408167432025920649097675171120131018942873360249725010269910294136950278971196796541847615455583289710059229092606788345933304998049916671909082483110382417806586624183591836284326050239004912679782227069703152518901877295727196972384011718334164197849058226419915338896604189779440276633", 10) 

	lcg := NewLCG(a, c, m, seed)
	result := lcg.Next(bits, false)
	
	if result.Cmp(&expectedResult) != 0 {
		t.Errorf("LCG.Next() = %v, want %v", result, &expectedResult)
	}
}

func BenchmarkLCG_Next(b *testing.B) {
	var seed uint64 = 16086232361280787691
	var m uint64 = 1 << 48 // 2⁴⁸
	var a uint64 = 25214903917
	var c uint64 = 11

	lcg := NewLCG(a, c, m, seed)

	for _, bits := range globals.TestSizes {	
		b.Run(fmt.Sprintf("bits %d", bits), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                lcg.Next(bits, false)
            }
        })
	}
}
