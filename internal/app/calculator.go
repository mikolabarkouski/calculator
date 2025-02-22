package app

import (
	"fmt"
	"math"
	"sort"
)

func (a *App) CalculatePacksNeeded(orderQuantity int, packSizes []int) (map[int]int, error) {
	if orderQuantity <= 0 {
		return nil, fmt.Errorf("order quantity must be a positive integer")
	}

	sort.Sort(sort.Reverse(sort.IntSlice(packSizes))) // Сортируем по убыванию

	dp := make([]int, orderQuantity+1)
	choice := make([]int, orderQuantity+1)

	for i := 1; i <= orderQuantity; i++ {
		dp[i] = math.MaxInt32 // Заполняем максимальными значениями
		for _, pack := range packSizes {
			if i >= pack && dp[i-pack]+1 < dp[i] {
				dp[i] = dp[i-pack] + 1
				choice[i] = pack
			}
		}
	}

	if dp[orderQuantity] == math.MaxInt32 {
		return nil, fmt.Errorf("cannot fulfill order with given pack sizes")
	}

	// Восстанавливаем ответ
	result := make(map[int]int)
	remaining := orderQuantity
	for remaining > 0 {
		pack := choice[remaining]
		result[pack]++
		remaining -= pack
	}

	return result, nil
}
