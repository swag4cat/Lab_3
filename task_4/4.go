package main

import (
	"bufio" // Для буферизованного чтения ввода
	"fmt"
	"math"
	"os" // Для работы с операционной системой (ввод/вывод)
)

func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите n и m: ")
	fmt.Fscan(reader, &n, &m)

	nums := make([]int, n)
	fmt.Print("Ввод: ")
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &nums[i])
	}

	// dp[i] — максимальная разница очков, которую может получить текущий игрок, начиная с позиции i
	dp := make([]int64, n+1)

	// suffix_sum[i] — сумма элементов от i до конца массива
	suffix_sum := make([]int64, n+1)

	for i := n - 1; i >= 0; i-- {
		suffix_sum[i] = suffix_sum[i+1] + int64(nums[i])
	}

	for i := n - 1; i >= 0; i-- {
		maxDiff := int64(math.MinInt64)

		// Игрок может взять от 1 до m чисел
		for j := 1; j <= m && i+j <= n; j++ {
			// Сумма взятых чисел: suffix_sum[i] - suffix_sum[i + j]
			// Разница очков: (взятая сумма) - (оптимальный результат противника)
			currentDiff := (suffix_sum[i] - suffix_sum[i+j]) - dp[i+j]
			if currentDiff > maxDiff {
				maxDiff = currentDiff
			}
		}

		dp[i] = maxDiff
	}

	// Если dp[0] > 0, Павел выиграл
	fmt.Print("Резульат: ")
	if dp[0] > 0 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
