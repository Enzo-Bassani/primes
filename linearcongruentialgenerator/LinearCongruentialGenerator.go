package linearcongruentialgenerator

import (
	"math/big"
)

// Classe do gerador.
type LCG struct {
	a                uint64
	c                uint64
	m                uint64
	x                uint64
	generatedNumbers []uint64 // Quando requisitado, é utilizado para testes
}

// Construtor da classe.
func NewLCG(a, c, m, seed uint64) *LCG {
	return &LCG{a, c, m, seed, []uint64{}}
}

// Next é a função principal do gerador. Retorna um número
// com tantos bits quanto requisitado. Se saveNumbers for verdadeiro,
// salva os números gerados para fins de teste.
func (lcg *LCG) Next(bits int, saveNumbers bool) *big.Int {
	// Divisão com arredondamento para cima
	requestedBytes := (bits + 7) / 8
	// Gera sempre um número com total de bytes múltiplo de 4
	totalBytes := requestedBytes
	if totalBytes % 4 != 0 {
		totalBytes = requestedBytes + 4 - requestedBytes % 4
	}

	if saveNumbers {
		lcg.generatedNumbers = []uint64{}
	}

	stream := make([]byte, totalBytes)
	bytesPerIteration := 4
	for i := 0; i < totalBytes; i += bytesPerIteration {
		// Gera o próximo número da sequência
		num := lcg.next()

		if saveNumbers {
			lcg.generatedNumbers = append(lcg.generatedNumbers, num)
		}

		// Extrai os bits 47 a 16
		putBits46To16(stream[i: i + bytesPerIteration], num)
	}

	// Descarta bytes extras
	stream = stream[:requestedBytes]
	var output big.Int
	return output.SetBytes(stream)
}

// Calcula o próximo valor de x.
// x = x² mod n
func (lcg *LCG) next() uint64 {
	lcg.x = (lcg.a*lcg.x + lcg.c) % lcg.m
	return lcg.x
}

func (lcg *LCG) GeneratedNumbers() []uint64 {
	return lcg.generatedNumbers
}

// Coloca em b os bits 47 a 16 de v.
func putBits46To16(b []byte, v uint64) {
	b[0] = byte(v >> 40)
	b[1] = byte(v >> 32)
	b[2] = byte(v >> 24)
	b[3] = byte(v >> 16)
}
