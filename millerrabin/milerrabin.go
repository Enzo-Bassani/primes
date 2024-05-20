package millerrabin

import (
	"crypto/rand"
	"math/big"
	g "primes/globals"
)

// ProbablyPrime retorna verdadeiro se n for aprovado
// pelo teste de Miller Rabin iterations vezes. Também
// retorna o número de iterações  que foram necessárias
// para verificar a primalidade.
func ProbablyPrime(n *big.Int, iterations int) (bool, int) {
	// Verifica se n é par
	if n.Bit(0) == 0 {
		return false, 0
	}

	// ---
	// 1. Escreva n - 1 = 2^k * m, onde m é ímpar
	// ---
	var left big.Int
	left.Sub(n, g.One)
	// Encontra o primeiro bit de valor 1
	k := 0
	for ; left.Bit(k) == 0; k++ {}

	// Right shift k vezes para obter m
	var m big.Int
	m.Rsh(&left, uint(k))

	// ---
	// 2. Executa o teste de Miller Rabin iterations vezes
	// ---
	for i := 0; i < iterations; i++ {
		if !millerRabinLoop(n, &m, k) {
			return false, i+1
		}
	}

	return true, iterations
}

// millerRabinLoop é a parte do teste que
// deve ser executada múltiplas vezes. Retorna
// verdadeiro se é provavelmente primo.
func millerRabinLoop(n, m *big.Int, k int) bool {
	a, _ := rand.Int(rand.Reader, n)

	var b big.Int
	b.Exp(a, m, n)

	var mod big.Int
	// if b ≡ 1 (mod n)
	if b.Cmp(g.One) == 0 {
		return true
	}

	var n_minus_1 big.Int
	n_minus_1.Sub(n, g.One)
	for i := 0; i < k; i++ {
		// Se b ≡ -1 (mod n)
		if mod.Mod(&b, n).Cmp(&n_minus_1) == 0 {
			return true
		} else {
			b.Exp(&b, g.Two, n)
		}
	}

	return false
}
