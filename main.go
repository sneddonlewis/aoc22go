package main

import "fmt"

func sumFood(food []int) (sum int) {
	for _, v := range food {
		sum += v
	}
	return sum
}

func calculateCalorificElf(elves [][]int) (ordinal int, calories int) {
	ordinal = -1
	calories = -1
	for i, food := range elves {
		currentCal := sumFood(food)
		if currentCal > calories {
			calories = currentCal
			ordinal = i + 1
		}
	}
	return ordinal, calories
}

func solve(elves ...[]int) {
	ordinal, calories := calculateCalorificElf(elves)
	fmt.Printf("Elf number %d carries the most calories at %d cal\n", ordinal, calories)
}

func main() {
	fmt.Println("Advent of Code 2022 in Go!")
	fmt.Println("Day 1")

	solve(
		[]int{
			1000,
			2000,
			3000,
		},
		[]int{
			4000,
		},
		[]int{
			5000,
			6000,
		},
		[]int{
			7000,
			8000,
			9000,
		},
		[]int{
			10000,
		},
	)
}
