package main

import (
	"fmt"
	"math"
)

// Проверка сходимости ряда по признаку Даламбера
func isConvergent(b int) bool {
	// Если 1/b < 1, т.е. b > 1, значит, ряд сходится
	return b > 1
}

// Вычисление НОД для сокращения дроби
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Попытка вычислить сумму ряда и вернуть её как несократимую дробь
func getRationalSum(a, b int) (int, int) {
	const EPS = 1e-10       // точность для прекращения суммирования
	const MAX_DENOMINATOR = 1000000 // для преобразования в дробь

	sum := 0.0
	for n := 1; ; n++ {
		term := math.Pow(float64(n), float64(a)) / math.Pow(float64(b), float64(n))
		if term < EPS {
			break
		}
		sum += term
	}

	// Преобразуем в несократимую дробь
	numerator := int(math.Round(sum * MAX_DENOMINATOR))
	gcdVal := gcd(numerator, MAX_DENOMINATOR)
	return numerator / gcdVal, MAX_DENOMINATOR / gcdVal
}

func main() {
	var a, b int
	fmt.Print("Введите значение для a: ")
	fmt.Scan(&a)
	fmt.Print("Введите значение для b: ")
	fmt.Scan(&b)

	// Если b == 1, тогда знаменатель не убывает, и ряд расходится
	if b == 1 {
		fmt.Println("Результат: infinity")
		return
	}

	// Проверка сходимости по признаку Даламбера
	if !isConvergent(b) {
		fmt.Println("Результат: infinity")
		return
	}

	// Вычисление суммы ряда
	numerator, denominator := getRationalSum(a, b)

	// Проверка на иррациональность: сравнение приближения с точной суммой
	approx := float64(numerator) / float64(denominator)
	exact := 0.0
	for n := 1; n < 100; n++ {
		exact += math.Pow(float64(n), float64(a)) / math.Pow(float64(b), float64(n))
	}

	if math.Abs(approx-exact) > 1e-6 {
		fmt.Println("Результат: irrational")
	} else {
		fmt.Printf("Результат: %d/%d\n", numerator, denominator)
	}
}
