package main

import (
	"fmt"
	"math"
)

// Функция, реализующая график, заданный по частям
func f(x float64) float64 {
	switch {
		case x >= -4 && x <= 0:
			// Полукруг: y = sqrt(4 - (x + 2)^2), центр в (-2, 0), радиус 2
			underRoot := 4 - math.Pow(x+2, 2)
			if underRoot >= 0 {
				return math.Sqrt(underRoot)
			}
			return math.NaN()
		case x > 0 && x < 2:
			// Функция: y = ln(x) / x
			return math.Log(x) / x
		case x >= 2 && x <= 4:
			// Постоянная функция: y = 1
			return 1
		default:
			// Вне области определения
			return math.NaN()
	}
}

func main() {
	// Интервал и шаг
	xStart := -4.0
	xEnd := 4.0
	dx := 0.5

	// Заголовок таблицы
	fmt.Println("\n╔════════════════════════════════════════╗")
	fmt.Println("║        ТАБЛИЦА ЗНАЧЕНИЙ ФУНКЦИИ        ║")
	fmt.Println("╠═══════════════╦════════════════════════╣")
	fmt.Println("║   Значение x  ║     Значение f(x)      ║")
	fmt.Println("╠═══════════════╬════════════════════════╣")

	// Основной цикл по X от xStart до xEnd с шагом dx
	for x := xStart; x <= xEnd+1e-6; x += dx {
		y := f(x) // вычисляем значение функции
		fmt.Printf("║%9.2f      ║ ", x)
		if !math.IsNaN(y) {
			// Если значение определено — выводим
			fmt.Printf("%15.5f        ║\n", y)
		} else {
			// Если значение не определено — выводим соответствующее сообщение
			fmt.Printf("%10s        ║\n", "нет значения")
		}
	}
	fmt.Println("╚═══════════════╩════════════════════════╝")
}
