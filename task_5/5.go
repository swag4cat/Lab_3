package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Функция расчёта температуры кофе по закону охлаждения Ньютона
func calculateCoffeeTemp(initialTemp, ambientTemp, coolingRate, time float64) float64 {
	return ambientTemp + (initialTemp-ambientTemp)*math.Exp(-coolingRate*time)
}

// Функция линейной аппроксимации по методу наименьших квадратов
func approximateLine(x, y []float64) (a, b float64) {
	n := len(x)
	var sumX, sumY, sumXY, sumX2 float64

	for i := 0; i < n; i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
	}

	a = (float64(n)*sumXY - sumX*sumY) / (float64(n)*sumX2 - sumX*sumX)
	b = (sumY - a*sumX) / float64(n)
	return a, b
}

// Функция вычисления коэффициента корреляции Пирсона
func correlation(x, y []float64) float64 {
	n := len(x)
	if n == 0 {
		return 0
	}

	var sumX, sumY float64
	for i := 0; i < n; i++ {
		sumX += x[i]
		sumY += y[i]
	}

	meanX := sumX / float64(n)
	meanY := sumY / float64(n)

	var numerator, sumSqX, sumSqY float64

	for i := 0; i < n; i++ {
		dx := x[i] - meanX
		dy := y[i] - meanY
		numerator += dx * dy
		sumSqX += dx * dx
		sumSqY += dy * dy
	}

	if sumSqX == 0 || sumSqY == 0 {
		return 0 // защита от деления на 0
	}

	return numerator / math.Sqrt(sumSqX*sumSqY)
}

// Функция для безопасного ввода положительного числа
func inputPositive(prompt string) float64 {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		value, err := strconv.ParseFloat(input, 64)
		if err != nil || value <= 0 {
			fmt.Println("Некорректный ввод. Повторите попытку.")
			continue
		}
		return value
	}
}

func main() {
	// Объявление переменных
	var times, temps []float64
	var initialTemp, ambientTemp, coolingRate, maxTime, timeStep float64

	// Вывод заголовка
	fmt.Println("╔════════════════════════════════╗")
	fmt.Println("║  МОДЕЛИРОВАНИЕ ОСТЫВАНИЯ КОФЕ  ║")
	fmt.Println("╚════════════════════════════════╝")

	// Ввод начальной температуры кофе
	fmt.Print("Введите начальную температуру кофе (в °C): ")
	fmt.Scan(&initialTemp)

	// Ввод температуры окружающей среды
	fmt.Print("Введите температуру окружающей среды (в °C): ")
	fmt.Scan(&ambientTemp)

	// Ввод коэффициента остывания (он должен быть положительным)
	coolingRate = inputPositive("Введите коэффициент остывания (в 1/мин): ")

	// Ввод общего времени моделирования
	maxTime = inputPositive("Введите общее время моделирования (в мин): ")

	// Ввод шага по времени
	timeStep = inputPositive("Введите шаг по времени (в мин): ")

	// Заголовок таблицы с результатами
	fmt.Println("\n╔════════════════════════════════════════╗")
	fmt.Println("║        РЕЗУЛЬТАТЫ МОДЕЛИРОВАНИЯ        ║")
	fmt.Println("╠═══════════════╦════════════════════════╣")
	fmt.Println("║   Время (мин) ║   Температура (°C)     ║")
	fmt.Println("╠═══════════════╬════════════════════════╣")

	// Цикл моделирования остывания в заданном промежутке времени
	for t := 0.0; t <= maxTime; t += timeStep {
		// Вычисляем текущую температуру с использованием функции
		currentTemp := calculateCoffeeTemp(initialTemp, ambientTemp, coolingRate, t)
		times = append(times, t)
		temps = append(temps, currentTemp)
		fmt.Printf("║ %13f ║ %18f     ║\n", t, currentTemp) // Выводим результаты в таблице
	}

	// Заключительный вывод таблицы
	fmt.Println("╚═══════════════╩════════════════════════╝")

	a, b := approximateLine(times, temps)

	// Выводим значения аппроксимирующей прямой
	fmt.Printf("\nАппроксимирующая прямая: y = %.4f * x + %.4f\n", a, b)

	r := correlation(times, temps)

	// Выводим значение коэффициента корреляции
	fmt.Printf("Коэффициент корреляции: r = %.4f\n", r)
}
