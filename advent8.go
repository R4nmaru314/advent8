package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const file = "input.txt"

func main() {
	file, _ := os.Open(file)
	scanner := bufio.NewScanner(file)

	var grid = make(map[int][]int)
	var outsideTrees int
	var insideTrees int
	var maxScenicScore int

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, "")
		for i, val := range values {
			num, err := strconv.Atoi(val)
			if err == nil {
				grid[i] = append(grid[i], num)
			}
		}
	}
	calculateOutsideTrees(grid, &outsideTrees)
	calculateInsideTreesAndMaxScenicScore(grid, &insideTrees, &maxScenicScore)
	log.Printf("Trees Outside: %d , Trees Inside: %d", outsideTrees, insideTrees)
	log.Printf("total: %d", outsideTrees+insideTrees)
	log.Printf("maxScenicScore: %d", maxScenicScore)
}

func calculateOutsideTrees(grid map[int][]int, outsideTrees *int) {
	*outsideTrees = 2*len(grid) + 2*len(grid[0]) - 4
}

func calculateInsideTreesAndMaxScenicScore(grid map[int][]int, insideTrees *int, maxScenicScore *int) {
	for height := 1; height < len(grid)-1; height++ {
		for length := 1; length < len(grid[0])-1; length++ {
			treeToCheck := grid[height][length]
			calculatePart1(grid, height, length, treeToCheck, insideTrees)
			calculatePart2(grid, height, length, treeToCheck, maxScenicScore)
		}
	}
}

func calculatePart1(grid map[int][]int, height int, length int, treeToCheck int, insideTrees *int) {
	if calculatePart1Left(grid, height, length, treeToCheck) {
		*insideTrees++
	} else if calculatePart1Right(grid, height, length, treeToCheck) {
		*insideTrees++
	} else if calculatePart1Top(grid, height, length, treeToCheck) {
		*insideTrees++
	} else if calculatePart1Bottom(grid, height, length, treeToCheck) {
		*insideTrees++
	}
}

func calculatePart2(grid map[int][]int, height int, length int, tree int, maxScenicScore *int) {
	left := calculatePart2Left(grid, height, length, tree)
	right := calculatePart2Right(grid, height, length, tree)
	top := calculatePart2Top(grid, height, length, tree)
	bottom := calculatePart2Bottom(grid, height, length, tree)
	scenicScore := left * right * top * bottom
	if scenicScore > *maxScenicScore {
		*maxScenicScore = scenicScore
	}
}

func calculateLeft(grid map[int][]int, height int, length int) []int {
	var left []int
	for i := 0; i < length; i++ {
		left = append(left, grid[height][i])
	}
	return left
}

func calculateRight(grid map[int][]int, height int, length int) []int {
	var right []int
	for i := len(grid[height]); i > length+1; i-- {
		right = append(right, grid[height][i-1])
	}
	return right
}

func calculateTop(grid map[int][]int, height int, length int) []int {
	var top []int
	for i := 0; i < height; i++ {
		top = append(top, grid[i][length])
	}
	return top
}

func calculateBottom(grid map[int][]int, height int, length int) []int {
	var bottom []int
	for i := len(grid); i > height+1; i-- {
		bottom = append(bottom, grid[i-1][length])
	}
	return bottom
}

func calculatePart1Left(grid map[int][]int, height int, length int, tree int) bool {
	return calculateLinePart1(calculateLeft(grid, height, length), tree)
}

func calculatePart1Right(grid map[int][]int, height int, length int, tree int) bool {
	return calculateLinePart1(calculateRight(grid, height, length), tree)
}

func calculatePart1Top(grid map[int][]int, height int, length int, tree int) bool {
	return calculateLinePart1(calculateTop(grid, height, length), tree)
}

func calculatePart1Bottom(grid map[int][]int, height int, length int, tree int) bool {
	return calculateLinePart1(calculateBottom(grid, height, length), tree)
}

func calculatePart2Left(grid map[int][]int, height int, length int, tree int) int {
	left := calculateLeft(grid, height, length)
	reverseIntArrayPart2(left)
	return calculateLinePart2(left, tree)
}

func calculatePart2Right(grid map[int][]int, height int, length int, tree int) int {
	right := calculateRight(grid, height, length)
	reverseIntArrayPart2(right)
	return calculateLinePart2(right, tree)
}

func calculatePart2Top(grid map[int][]int, height int, length int, tree int) int {
	top := calculateTop(grid, height, length)
	reverseIntArrayPart2(top)
	return calculateLinePart2(top, tree)
}

func calculatePart2Bottom(grid map[int][]int, height int, length int, tree int) int {
	bottom := calculateBottom(grid, height, length)
	reverseIntArrayPart2(bottom)
	return calculateLinePart2(bottom, tree)
}

func calculateLinePart1(array []int, tree int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] >= tree {
			return false
		}
	}
	return true
}

func calculateLinePart2(array []int, tree int) int {
	for i := 0; i < len(array); i++ {
		if array[i] >= tree {
			return i + 1
		}
	}
	return len(array)
}

func reverseIntArrayPart2(array []int) {
	length := len(array)
	for i := 0; i < length/2; i++ {
		array[i], array[length-i-1] = array[length-i-1], array[i]
	}
}
