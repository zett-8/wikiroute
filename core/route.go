package core

import (
	"fmt"
	"sort"
)

func sortByHeuristic(nodes []int, goalID int) {
	sort.Slice(nodes, func(i, j int) bool {
		return abs(nodes[i]-goalID) < abs(nodes[j]-goalID)
	})
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func PageExists(pageID int, pagesMap *PageData, pageLinksMap *PageLinksData) bool {
	if _, exists := pagesMap.IDToTitle[int32(pageID)]; exists {
		return true
	}
	return false
}

func BidirectionalBFS(fromID, toID int, pageLinksMap *PageLinksData) ([]int, bool, int) {
	if fromID == toID {
		return []int{fromID}, true, 0
	}

	readCount := 0
	maxSteps := 3
	forwardQueue := [][]int{{fromID}}
	backwardQueue := [][]int{{toID}}
	forwardVisited := map[int][]int{fromID: {fromID}}
	backwardVisited := map[int][]int{toID: {toID}}

	for step := 0; step < maxSteps; step++ {
		newForwardQueue := [][]int{}
		for _, path := range forwardQueue {
			current := path[len(path)-1]
			links := GetLinksFromDB(current, &readCount, pageLinksMap)
			sortByHeuristic(links, toID)
			for _, next := range links {
				if _, exists := forwardVisited[next]; exists {
					continue
				}
				newPath := append([]int{}, path...)
				newPath = append(newPath, next)
				forwardVisited[next] = newPath
				if backwardPath, exists := backwardVisited[next]; exists {
					fullPath := append(newPath[:len(newPath)-1], ReversePath(backwardPath)...)
					return fullPath, true, readCount
				}
				newForwardQueue = append(newForwardQueue, newPath)
			}
		}
		forwardQueue = newForwardQueue

		newBackwardQueue := [][]int{}
		for _, path := range backwardQueue {
			current := path[len(path)-1]
			links := GetLinksToDB(current, &readCount, pageLinksMap)
			sortByHeuristic(links, fromID)
			for _, prev := range links {
				if _, exists := backwardVisited[prev]; exists {
					continue
				}
				newPath := append([]int{}, path...)
				newPath = append(newPath, prev)
				backwardVisited[prev] = newPath
				if forwardPath, exists := forwardVisited[prev]; exists {
					fullPath := append(forwardPath[:len(forwardPath)-1], ReversePath(newPath)...)
					return fullPath, true, readCount
				}
				newBackwardQueue = append(newBackwardQueue, newPath)
			}
		}
		backwardQueue = newBackwardQueue
	}

	return nil, false, readCount
}

func GetLinksFromDB(pageID int, readCount *int, pageLinksMap *PageLinksData) []int {
	*readCount++
	if links, exists := pageLinksMap.PageLinksMap[int32(pageID)]; exists {
		var result []int
		for _, link := range links {
			result = append(result, int(link))
		}
		return result
	}
	return nil
}

func GetLinksToDB(pageID int, readCount *int, pageLinksMap *PageLinksData) []int {
	*readCount++
	if links, exists := pageLinksMap.PageLinksMapReverse[int32(pageID)]; exists {
		var result []int
		for _, link := range links {
			result = append(result, int(link))
		}
		return result
	}
	return nil
}

func GetPageIDByTitle(title string, pagesMap *PageData) (int, error) {
	if id, exists := pagesMap.TitleToID[title]; exists {
		return int(id), nil
	}
	return 0, fmt.Errorf("page not found")
}

func GetPageTitleByID(id int, pagesMap *PageData) (string, error) {
	if title, exists := pagesMap.IDToTitle[int32(id)]; exists {
		return title, nil
	}
	return "", fmt.Errorf("page not found")
}

func ReversePath(path []int) []int {
	reversed := make([]int, len(path))
	for i := range path {
		reversed[i] = path[len(path)-1-i]
	}
	return reversed
}
