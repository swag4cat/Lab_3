package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Функция для возведения числа в степень по модулю
func powMod(base, degree, mod int) int {
	result := base % mod
	// Возводим основание в степень по модулю с использованием метода быстрого возведения в степень
	for i := 2; i <= degree; i++ {
		result = (result * base) % mod
	}
	return result
}

// Решето Эратосфена для нахождения всех простых чисел до N
func sieveEratos(N int) []int {
	boolPrime := make([]bool, N+1)
	// Инициализируем все числа как простые
	for i := 2; i <= N; i++ {
		boolPrime[i] = true
	}

	// Основной цикл решета Эратосфена для исключения составных чисел
	for p := 2; p*p <= N; p++ {
		if boolPrime[p] {
			// Если число простое, помечаем все его кратные как составные
			for i := p * p; i <= N; i += p {
				boolPrime[i] = false
			}
		}
	}

	// Формируем список всех простых чисел
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
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
	return rand.Intn(max-min+1) + min
}

// Тест Рабина-Миллера на простоту
func rabin(num, k int) bool {
	if num == 2 || num == 3 {
		return true // 2 и 3 - простые
	}
	if num < 2 || num%2 == 0 {
		return false // Четные числа (кроме 2) составные
	}

	// Разлагаем число num-1 как 2^s * d
	s := 0
	d := num - 1
	for d%2 == 0 {
		d /= 2
		s++
	}

	// Повторяем k раз для случайных значений a
	for i := 0; i < k; i++ {
		a := rand.Intn(num-4) + 2
		x := powMod(a, d, num) // a^d % num
		y := 0
		for j := 0; j < s; j++ {
			y = powMod(x, 2, num) // Возведение в квадрат
			if y == 1 && x != 1 && x != (num-1) {
				return false // Если x = 1 или x = num-1, продолжаем, иначе число составное
			}
			x = y
		}
		if y != 1 {
			return false // Если не получилось получить 1, число составное
		}
	}
	return true // Число вероятно простое
}

// Тест Поклингтона на простоту
func poklington(n, t int, qList []int) bool {
	aSet := make(map[int]bool)

	// Генерируем t случайных чисел a, где 2 <= a < n
	for len(aSet) < t {
		a := randDist(2, n-1)
		aSet[a] = true
	}

	// Проверяем условие для каждого a: a^(n-1) % n == 1
	for a := range aSet {
		if powMod(a, n-1, n) != 1 {
			return false
		}
	}

	// Проверяем условие для каждого a, чтобы ни одно из a^(n-1/q) не было равно 1 для q из qList
	for a := range aSet {
		allNotOne := true
		for _, q := range qList {
			if powMod(a, (n-1)/q, n) == 1 {
				allNotOne = false
				break
			}
		}
		if allNotOne {
			return true // Если все условия выполнены, число вероятно простое
		}
	}

	return false // Если не выполнены условия для любого a, число составное
}

// Генерация числа n и списка простых множителей n-1
func calcN(primes []int, bit int) (int, []int) {
	maxIndex := len(primes) - 1
	maxPow := 1
	for ; math.Pow(2, float64(maxPow)) < math.Pow(2, float64(bit/2+1)); maxPow++ {
	}

	min := bit/2 + 1
	max := bit/2 + 2
	powMax := math.Pow(2, float64(max))

	f := 1
	qList := []int{}

	// Генерация n как произведения простых чисел
	for {
		randIndex := randDist(0, maxIndex)
		alphaDegree := randDist(1, maxPow)

		if math.Pow(float64(primes[randIndex]), float64(alphaDegree)) > 0 {
			if float64(f)*math.Pow(float64(primes[randIndex]), float64(alphaDegree)) < powMax {
				f *= int(math.Pow(float64(primes[randIndex]), float64(alphaDegree)))
				qList = append(qList, primes[randIndex])
			}
		}

		// Проверка на достижение нужного диапазона для числа
		if float64(f) > math.Pow(2, float64(bit/2)) {
			if float64(f) >= math.Pow(2, float64(bit/2+1)) {
				f = 1
				qList = []int{}
			} else {
				break
			}
		}
	}

	// Генерация случайного числа R
	R := randDist(min-1, max-1)
	if R%2 != 0 {
		R++ // Делаем R четным
	}
	n := R*f + 1 // Получаем число n

	// Удаляем дубликаты из списка простых множителей
	unique := make(map[int]bool)
	var uniqueQList []int
	for _, q := range qList {
		if !unique[q] {
			unique[q] = true
			uniqueQList = append(uniqueQList, q)
		}
	}

	return n, uniqueQList
}

func main() {
	bit := 10 // Битовая длина числа
	t := 10 // Количество тестов для Поклингтона
	primes := sieveEratos(500) // Получаем простые числа до 500

	var result []int // Список найденных простых чисел
	var testVer []string // Результаты теста Рабина-Миллера
	var kCont []int // Счётчик числа успешных тестов

	k := 0
	// Ищем 10 вероятно простых чисел
	for len(result) < 10 {
		n, qList := calcN(primes, bit) // Генерация числа n

		// Проверяем число с помощью теста Поклингтона
		if poklington(n, t, qList) {
			result = append(result, n)

			// Проверяем число с помощью теста Рабина-Миллера
			if rabin(n, 3) {
				testVer = append(testVer, "+")
			} else {
				testVer = append(testVer, "-")
			}
			kCont = append(kCont, k)
			k = 0
		} else {
			if rabin(n, 3) {
				k++ // Увеличиваем счётчик, если тест Рабина-Миллера прошёл
			}
		}
	}

	// Вывод результатов в виде таблицы
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
