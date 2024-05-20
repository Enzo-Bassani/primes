package blumblumshub

import (
	"crypto/rand"
	"log"
	"math/big"
	g "primes/globals"
)

// Classe do gerador. Guarda os valores de
// N e o valor atual de X.
type BBS struct {
	n *big.Int
	x *big.Int
}

// Construtor da classe
func NewBBS(seed, p, q *big.Int) *BBS {
	n := big.NewInt(0)

	n = n.Mul(p, q)
	// seed ∈ R [1, n − 1]
	if seed.Cmp(g.One) == -1 || seed.Cmp(n) == 1 {
		log.Fatal("a seed deve estar entre 1 e p * q")
	}

	if seed.Cmp(p) == 0 || seed.Cmp(q) == 0 {
		log.Fatal("A seed deve ser um co-primo de N")
	}

	bbs := &BBS{
		n: n,
		x: seed,
	}

	// Realiza a primeira iteração do algoritmo.
	bbs.next()

	return bbs
}

// Gera um primo Blum.
func GetBlumPrime() *big.Int {
	mod := big.NewInt(0)

	// Gera primos aleatórios até obter um congruente a 3 mod 4.
	p, _ := rand.Prime(rand.Reader, 128)
	for mod.Mod(p, g.Four).Cmp(g.Three) != 0 {
		p, _ = rand.Prime(rand.Reader, 128)
	}

	return p
}

// Next é a função principal do gerador. Retorna um número
// com tantos bits quanto requisitado.
func (bbs *BBS) Next(bits int) *big.Int {
	bytes := (bits + 7) / 8
	stream := make([]byte, bytes)

	// Para cada byte
	for i := 0; i < bytes; i++ {
		// Para cada bit
		for j := 0; j < 8; j++ {
			// Gera o próximo número da sequência
			bbs.next()
			// Extrai o LSB
			stream[i] |= byte(bbs.x.Bit(0)) << (7 - j)
		}
	}

	// Transforma o array de bytes em uma variável númerica
	var output big.Int
	return output.SetBytes(stream)
}

// Calcula o próximo valor de x.
// x = x² mod n
func (bbs *BBS) next() {
	bbs.x.Exp(bbs.x, g.Two, bbs.n)
}
