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
			left, right, top, bottom := calculateSides(grid, height, length)
			calculatePart1(left, right, top, bottom, treeToCheck, insideTrees)
			calculatePart2(left, right, top, bottom, treeToCheck, maxScenicScore)
		}
	}
}

func calculatePart1(left, right, top, bottom []int, tree int, insideTrees *int) {
	if calculateLinePart1(left, tree) {
		*insideTrees++
	} else if calculateLinePart1(right, tree) {
		*insideTrees++
	} else if calculateLinePart1(top, tree) {
		*insideTrees++
	} else if calculateLinePart1(bottom, tree) {
		*insideTrees++
	}
}

func calculatePart2(left, right, top, bottom []int, tree int, maxScenicScore *int) {
	leftScore := calculateLinePart2(reverseArrayPart2(left), tree)
	rightScore := calculateLinePart2(reverseArrayPart2(right), tree)
	topScore := calculateLinePart2(reverseArrayPart2(top), tree)
	bottomScore := calculateLinePart2(reverseArrayPart2(bottom), tree)
	scenicScore := leftScore * rightScore * topScore * bottomScore
	if scenicScore > *maxScenicScore {
		*maxScenicScore = scenicScore
	}
}

func calculateSides(grid map[int][]int, height int, length int) ([]int, []int, []int, []int) {
	left := calculateLeft(grid, height, length)
	right := calculateRight(grid, height, length)
	top := calculateTop(grid, height, length)
	bottom := calculateBottom(grid, height, length)

	return left, right, top, bottom
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

func reverseArrayPart2(array []int) []int {
    length := len(array)
    reversedArray := make([]int, length)
    for i := 0; i < length; i++ {
        reversedArray[i] = array[length-i-1]
    }
    return reversedArray
}
