package main

import (
	"sort"
)

type Cluster struct {
	buildings  [][]int
	points     map[int]int
	rectangles [][]int
}

func constructor(building []int) *Cluster {
	buildings := [][]int{building}
	points := map[int]int{
		building[0]: building[2],
	}
	rectangles := [][]int{building}
	return &Cluster{buildings, points, rectangles}
}

func (cluster *Cluster) genContour() {
	for i := 1; i < len(cluster.buildings); i++ {
		building := cluster.buildings[i]
		start, end := cluster.getAffectedRectangle(building)
		cluster.affectFromStart(building, start, end)
	}
}

func (cluster *Cluster) affectFromStart(building []int, start int, end int) {
	rectangle := cluster.rectangles[start]
	if building[2] > rectangle[2] {
		cluster.insertPoint(building[0], rectangle[2])
		realEnd := end
		if realEnd == -1 {
			realEnd = len(cluster.rectangles) - 1
		}
		for i := start + 1; i <= realEnd; i++ {
			cluster.removePoint(cluster.rectangles[i][0])
		}
		newRectangles := [][]int{building}
		if end != -1 {
			newRectangles = append(newRectangles, []int{building[1], cluster.rectangles[end][1], cluster.rectangles[end][2]})
			for i := end + 1; i < len(cluster.rectangles); i++ {
				newRectangles = append(newRectangles, cluster.rectangles[i])
			}
		}
		cluster.rectangles = newRectangles
	} else {
		if end == -1 {
			end = len(cluster.rectangles) - 1
		}
		crossIndex := -1
		for i := start + 1; i <= end; i++ {
			if crossIndex == -1 && cluster.rectangles[i][2] >= building[2] {
				cluster.insertPoint(cluster.rectangles[i][0], building[2])
				crossIndex = i
			}
			if crossIndex != -1 && cluster.rectangles[i][2] >= building[2] {
				cluster.removePoint(cluster.rectangles[i][0])
			}
		}
	}
}

func (cluster *Cluster) getAffectedRectangle(building []int) (start int, end int) {
	start = -1
	end = -1
	for i := 0; i < len(cluster.rectangles); i++ {
		rectangle := cluster.rectangles[i]
		if building[0] >= rectangle[0] && building[0] <= rectangle[1] {
			start = i
		}
		if building[1] >= rectangle[0] && building[1] <= rectangle[1] {
			end = i
			return
		}
	}
}

func (cluster *Cluster) insertPoint(x int, y int) {
	y1, ok := cluster.points[x]
	if ok {
		if y > y1 {
			cluster.points[x] = y
		}
	} else {
		cluster.points[x] = y
	}
}

func (cluster *Cluster) removePoint(x int) {
	delete(cluster.points, x)
}

func (cluster *Cluster) getContour() [][]int {
	contour := make([][]int, 0)
	for x, y := range cluster.points {
		contour = append(contour, []int{x, y})
	}
	return contour
}

func (cluster *Cluster) getRightest() []int {
	x := cluster.buildings[len(cluster.buildings)-1][1]
	return []int{x, 0}
}

func getSkyline(buildings [][]int) [][]int {
	clusters := split(buildings)
	for _, cluster := range clusters {
		cluster.genContour()
	}
	return getContour(clusters)
}

func getContour(clusters []*Cluster) [][]int {
	if len(clusters) == 0 {
		return make([][]int, 0)
	}
	results := make([][]int, 0)
	for _, cluster := range clusters {
		results = append(results, cluster.getContour()...)
	}
	results = append(results, clusters[len(clusters)-1].getRightest())
	return results
}

func split(buildings [][]int) []*Cluster {
	sort.SliceStable(buildings, func(i, j int) bool {
		b1 := buildings[i]
		b2 := buildings[j]
		if b1[0] < b2[0] {
			return true
		} else if b1[0] > b2[0] {
			return false
		}
		if b1[1] < b2[1] {
			return true
		} else if b1[1] > b2[1] {
			return false
		}
		if b1[2] < b2[2] {
			return false
		} else if b1[2] > b2[2] {
			return true
		}
		return false
	})
	clusters := make([]*Cluster, 0)
	if len(buildings) == 0 {
		return clusters
	}
	cur := constructor(buildings[0])
	clusters = append(clusters, cur)
	for i := 1; i < len(buildings); i++ {
		building := buildings[i]
		if building[0] <= cur.buildings[len(cur.buildings)-1][1] {
			cur.buildings = append(cur.buildings, building)
		} else {
			cur = constructor(building)
			clusters = append(clusters, cur)
		}
	}
	return clusters
}

func main() {
	buildings := [][]int{
		[]int{2, 9, 10},
		[]int{3, 7, 15},
		[]int{5, 12, 12},
		[]int{15, 20, 10},
		[]int{19, 24, 8},
	}
	getSkyline(buildings)
}
