package main

import (
	"BungieARG/database"
	"fmt"
	"os"
)

type Coord struct {
	X, Y int
}

func main() {

	database.InitDB()
	if os.Args[1] == "data" {
		database.LoadData()
	} else {
		Scout(os.Args[1])
	}

}

func Scout(startingFreq string) {
	limiter := "RRRRRRRR"

	visited := make(map[string]bool)
	queue := []string{startingFreq}

	positions := make(map[string]Coord)

	positions[startingFreq] = Coord{X: 0, Y: 0}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}

		visited[current] = true
		pos := positions[current]

		// Get directions (not freqs yet)
		Top, Bottom, Left, Right := database.FetchRelations(current)

		// Map of directions to handlers
		dirs := map[string]struct {
			offsetX, offsetY int
			fetchFunc        func(string) string
		}{
			Top:    {0, -1, database.FetchFreqTop},
			Bottom: {0, 1, database.FetchFreqBottom},
			Left:   {-1, 0, database.FetchFreqLeft},
			Right:  {1, 0, database.FetchFreqRight},
		}

		for dir, meta := range dirs {
			if dir != "" && dir != limiter {
				nextFreq := meta.fetchFunc(dir)
				if nextFreq != "" && !visited[nextFreq] {
					positions[nextFreq] = Coord{pos.X + meta.offsetX, pos.Y + meta.offsetY}
					queue = append(queue, nextFreq)
				}
			}
		}
	}

	printMap(positions)
	printMapPieces(positions)
}

func printMap(positions map[string]Coord) {
	// Find min/max bounds
	minX, maxX := 0, 0
	minY, maxY := 0, 0

	for _, pos := range positions {
		if pos.X < minX {
			minX = pos.X
		}
		if pos.X > maxX {
			maxX = pos.X
		}
		if pos.Y < minY {
			minY = pos.Y
		}
		if pos.Y > maxY {
			maxY = pos.Y
		}
	}

	// Invert Y for top-down printing
	grid := make(map[int]map[int]string)
	for freq, pos := range positions {
		if _, exists := grid[pos.Y]; !exists {
			grid[pos.Y] = make(map[int]string)
		}
		grid[pos.Y][pos.X] = freq
	}

	// Print grid
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if row, exists := grid[y]; exists {
				if val, ok := row[x]; ok {
					fmt.Printf("[%s]", val)
				} else {
					fmt.Print("[    ]")
				}
			} else {
				fmt.Print("[    ]")
			}
		}
		fmt.Println()
	}
}

func printMapPieces(positions map[string]Coord) {
	// Find min/max bounds
	minX, maxX := 0, 0
	minY, maxY := 0, 0

	// Create or open a file
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Ensure file is closed when done

	for _, pos := range positions {
		if pos.X < minX {
			minX = pos.X
		}
		if pos.X > maxX {
			maxX = pos.X
		}
		if pos.Y < minY {
			minY = pos.Y
		}
		if pos.Y > maxY {
			maxY = pos.Y
		}
	}

	// Invert Y for top-down printing
	grid := make(map[int]map[int]string)
	for freq, pos := range positions {
		if _, exists := grid[pos.Y]; !exists {
			grid[pos.Y] = make(map[int]string)
		}
		grid[pos.Y][pos.X] = freq
	}

	// Print grid
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if row, exists := grid[y]; exists {
				if val, ok := row[x]; ok {
					piece := database.FetchMidPiece(val)
					_, err = file.WriteString(fmt.Sprintf("[%s]", piece))
				} else {
					_, err = file.WriteString(fmt.Sprint("[    ]"))

				}
			} else {
				_, err = file.WriteString(fmt.Sprint("[    ]"))

			}
		}
		fmt.Println()
	}
}
