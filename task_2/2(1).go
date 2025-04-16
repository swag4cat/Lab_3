package main

import (
	"fmt"       // Для форматированного ввода/вывода
	"math"      // Математические функции
	"math/rand" // Генерация случайных чисел
	"time"      // Работа со временем
)

// Функция для возведения числа в степень по модулю
func powMod(base, degree, mod int) int {
	result := base % mod
	for i := 2; i <= degree; i++ {
		result = (result * base) % mod
	}
	return result
}

// Решето Эратосфена - алгоритм нахождения всех простых чисел до N
func sieveEratos(N int) []int {
	// Создаем массив флагов: true - простое, false - составное
	boolPrime := make([]bool, N+1)
	for i := 2; i <= N; i++ {
		boolPrime[i] = true
	}

	// Просеиваем: для каждого простого p вычеркиваем числа p*p
	for p := 2; p*p <= N; p++ {
		if boolPrime[p] {
			for i := p * p; i <= N; i += p {
				boolPrime[i] = false
			}
		}
	}

	// Собираем все простые числа
	primes := []int{}
	for p := 2; p <= N; p++ {
		if boolPrime[p] {
			primes = append(primes, p)
		}
	}

	return primes
}

// Генерация случайного числа в диапазоне [min, max]
func randDist(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// Тест Рабина-Миллера на простоту числа num с k раундами проверки
func rabin(num, k int) bool {
	if num == 2 || num == 3 {
		return true
	}
	if num < 2 || num%2 == 0 {
		return false
	}

	// Представляем num-1 в виде d*2^s
	s := 0
	d := num - 1
	for d%2 == 0 {
		d /= 2
		s++
	}

	// Проводим k раундов проверки
	for i := 0; i < k; i++ {
		a := rand.Intn(num-4) + 2
		x := powMod(a, d, num)
		y := 0
		for j := 0; j < s; j++ {
			y = powMod(x, 2, num)
			if y == 1 && x != 1 && x != (num-1) {
				return false
			}
			x = y
		}
		if y != 1 {
			return false
		}
	}
	return true
}

// Генерация кандидата на простое число n и его простых множителей
func calcN(primes []int, bit int) (int, []int) {
	maxIndex := len(primes) - 1
	maxPow := 1
	// Находим максимальную степень, при которой 2^maxPow < 2^(bit-1)
	for ; math.Pow(2, float64(maxPow)) < math.Pow(2, float64(bit-1)); maxPow++ {
	}

	m := 1
	qList := []int{}

	// Генерируем m как произведение случайных простых чисел в случайных степенях
	for {
		randIndex := randDist(0, maxIndex)
		alphaDegree := randDist(1, maxPow)

		if math.Pow(float64(primes[randIndex]), float64(alphaDegree)) > 0 {
			if float64(m)*math.Pow(float64(primes[randIndex]), float64(alphaDegree)) < math.Pow(2, float64(bit-1)) {
				m *= int(math.Pow(float64(primes[randIndex]), float64(alphaDegree)))
				qList = append(qList, primes[randIndex])
			}
		}

		// Проверяем, попадает ли m в нужный диапазон
		if float64(m) > math.Pow(2, float64(bit-2)) {
			if float64(m) >= math.Pow(2, float64(bit-1)) {
				m = 1
				qList = []int{}
			} else {
				break
			}
		}
	}

	// Вычисляем n = 2*m + 1 (кандидат на простое число)
	n := 2*m + 1
	return n, qList
}

// Тест Миллера на простоту числа n с t случайными основаниями
func Miller(n, t int, qList []int) bool {
	// Множество случайных оснований
	aSet := make(map[int]bool)

	// Генерируем t различных оснований
	for len(aSet) < t {
		a := randDist(2, n-1)
		aSet[a] = true
	}

	// Проверяем малую теорему Ферма
	for a := range aSet {
		if powMod(a, n-1, n) != 1 {
			return false
		}
	}

	// Проверяем дополнительные условия для каждого простого делителя q
	for _, q := range qList {
		isPrime := true
		for a := range aSet {
			res := powMod(a, (n-1)/q, n)
			if res != 1 {
				isPrime = false
				break
			}
		}
		if isPrime {
			return false
		}
	}

	return true
}

func main() {
	bit := 10  // Битовый размер генерируемых чисел
	t := 5     // Количество проверок в тесте Миллера
	primes := sieveEratos(500)  // Получаем список простых чисел до 500

	var result []int    // Найденные "простые" числа
	var testVer []string  // Результаты проверки тестом Рабина-Миллера (+/-)
	var kCont []int     // Счетчик отвергнутых кандидатов
	k := 0

	// Генерируем 10 чисел, прошедших тест Миллера
	for len(result) < 10 {
		n, qList := calcN(primes, bit)

		if Miller(n, t, qList) {
			result = append(result, n)

			// Проверяем сгенерированное число тестом Рабина-Миллера
			if rabin(n, 3) {
				testVer = append(testVer, "+")
			} else {
				testVer = append(testVer, "-")
			}
			kCont = append(kCont, k)
			k = 0
		} else {
			if rabin(n, 3) {
				k++
			}
		}
	}

	// Вывод результатов
	fmt.Println("\n╔═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Print("║Num\t║")
	for i := 1; i < 11; i++ {
		fmt.Printf("\t%d\t║", i)
	}
	fmt.Println("\n╠═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")

	fmt.Print("║P\t║")
	for _, item := range result {
		fmt.Printf("\t%d\t║", item)
	}
	fmt.Println("\n╠═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")

	fmt.Print("║Test\t║")
	for _, item := range testVer {
		fmt.Printf("\t%s\t║", item)
	}
	fmt.Println("\n╠═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣")

	fmt.Print("║k\t║")
	for _, item := range kCont {
		fmt.Printf("\t%d\t║", item)
	}
	fmt.Println("\n╚═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
}
