package main

import (
	"fmt"
)

// FIFOPageReplacement implements the FIFO page replacement algorithm
func FIFOPageReplacement(pages []int, capacity int) int {
	var pageFaults int
	queue := make([]int, 0, capacity)
	pageSet := make(map[int]bool)

	for _, page := range pages {
		// Check if the page is not in the set
		if !pageSet[page] {
			pageFaults++

			// If the queue is full, remove the oldest page
			if len(queue) == capacity {
				removedPage := queue[0]
				queue = queue[1:]
				delete(pageSet, removedPage)
			}

			// Add the new page to the queue and set
			queue = append(queue, page)
			pageSet[page] = true
		}
	}

	return pageFaults
}

func main() {
	pages := []int{1, 3, 0, 3, 5, 6, 3}
	capacity := 3
	pageFaults := FIFOPageReplacement(pages, capacity)
	fmt.Printf("Number of page faults: %d\n", pageFaults)
}
