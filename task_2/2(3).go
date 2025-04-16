package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Быстрое возведение в степень по модулю
func powMod(base, exponent, mod int) int {
	result := 1
	base %= mod
	for exponent > 0 {
		if exponent&1 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exponent >>= 1
	}
	return result
}

// Генерация случайного целого числа [min, max]
func randInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// Генерация случайного вещественного числа ξ ∈ [0, 1)
func randDouble() float64 {
	return rand.Float64()
}

// Решето Эратосфена
func sieveEratos(N int) []int {
	isPrime := make([]bool, N+1)
	for i := range isPrime {
		isPrime[i] = true
	}
	primes := []int{}
	for p := 2; p <= N; p++ {
		if isPrime[p] {
			primes = append(primes, p)
			for i := p * 2; i <= N; i += p {
				isPrime[i] = false
			}
		}
	}
	return primes
}

// Тест Миллера-Рабина
func millerRabin(n, iterations int) bool {
	if n < 2 || n%2 == 0 {
		return false
	}
	d := n - 1
	r := 0
	for d%2 == 0 {
		d /= 2
		r++
	}

	for i := 0; i < iterations; i++ {
		a := randInt(2, n-2)
		x := powMod(a, d, n)
		if x == 1 || x == n-1 {
			continue
		}
		passed := false
		for j := 0; j < r-1; j++ {
			x = powMod(x, 2, n)
			if x == n-1 {
				passed = true
				break
			}
		}
		if !passed {
			return false
		}
	}
	return true
}

// Генерация простого числа p по ГОСТ Р 34.10-94 (ξ ≠ 0)
func generateGOSTPrime(primes []int, bitLen int) int {
	qBitLen := (bitLen + 1) / 2
	qMin := 0
	qMax := (1 << qBitLen) - 1

	var q int
	for {
		q = primes[randInt(0, len(primes)-1)]
		if q >= qMin && q <= qMax {
			break
		}
	}

	half := 1 << (bitLen - 1) // целое число
	for {
		ξ := randDouble()
		n := (float64(half) + ξ*float64(half)) / float64(q)
		nInt := int(n)
		if nInt%2 != 0 {
			nInt++
		}

		for k := 0; ; k += 2 {
			p := (nInt + k)*q + 1
			if p >= (1 << bitLen) {
				break
			}
			if powMod(2, p-1, p) == 1 && powMod(2, nInt+k, p) != 1 {
				return p
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	bitLength := 12
	testRounds := 5
	primesSet := make(map[int]bool)
	primes := sieveEratos(500)
	testResults := []string{}
	rejected := 0

	for len(primesSet) < 10 {
		p := generateGOSTPrime(primes, bitLength)
		if primesSet[p] {
			continue
		}
		if millerRabin(p, testRounds) {
			primesSet[p] = true
			testResults = append(testResults, "ПРОЙДЕН")
		} else {
			rejected++
		}
	}

	fmt.Println("╔═════╦═══════╦═══════════════════════╗")
	fmt.Println("║  №  ║   P   ║ Тест Миллера-Рабина   ║")
	fmt.Println("╠═════╬═══════╬═══════════════════════╣")

	i := 1
	for p := range primesSet {
		fmt.Printf("║ %3d ║ %5d ║ %21s ║\n", i, p, testResults[i-1])
		i++
	}

	fmt.Println("╚═════╩═══════╩═══════════════════════╝")
	fmt.Println("K (отвергнутых Миллером):", rejected)
}
