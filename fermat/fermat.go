package fermat

import (
	"crypto/rand"
	"math/big"
	g "primes/globals"
)

// ProbablyPrime retorna verdadeiro se n for aprovado
// pelo teste de Fermat iterations vezes. Também
// retorna o número de iterações  que foram necessárias
// para verificar a primalidade.
func ProbablyPrime(p *big.Int, iterations int) (bool, int) {
	// Verifica se p é par.
	if p.Bit(0) == 0 {
		return false, 0
	}

	var mod, p_minus_one, p_minus_three big.Int
	p_minus_one.Sub(p, g.One)
	p_minus_three.Sub(p, g.Three)
	// Para cada iteração requisitada
	for i := 0; i < iterations; i++ {
		// Escolhe um A aleatório em (1, p-1)
		a, _ := rand.Int(rand.Reader, &p_minus_three)
		a.Add(a, g.Two)

		// Testa se A é coprimo de P
		var coprime big.Int
		if coprime.GCD(nil, nil, a, p).Cmp(g.One) != 0 {
			// Se não é, então P não pode ser primo.
			return false, i+1
		}

		// Se a^p-1 !≡ 1 (mod p)
		if mod.Exp(a, &p_minus_one, p).Cmp(g.One) != 0 {
			return false, i+1
		}
	}

	return true, iterations
}
