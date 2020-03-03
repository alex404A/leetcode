package main

import (
	"fmt"
	"strconv"
)

type Point struct {
	x int
	y int
}

func (this *Point) buildKey() string {
	return strconv.Itoa(this.x) + "." + strconv.Itoa(this.y)
}

func (this *Point) slope(other *Point) string {
	if other.x == this.x {
		return "-0.0"
	}
	if other.y == this.y {
		return "+0.0"
	}
	a1 := float64(other.x-this.x) / float64(other.y-this.y)
	// b1 := (other.x - this.x) % (other.y - this.y)
	return strconv.FormatFloat(a1, 'f', 32, 64)
}

type Line struct {
	list []*Point
}

func (this *Line) length(repeatedMap *RepeatedMap) int {
	length := len(this.list)
	for _, p := range this.list {
		key := p.buildKey()
		times, _ := repeatedMap.m[key]
		if times > 1 {
			length += times - 1
		}
	}
	return length
}

func (this *Line) append(point *Point) {
	this.list = append(this.list, point)
}

func (this *Line) copy(include *Point, exclude *Point) *Line {
	other := make([]*Point, 0)
	for _, point := range this.list {
		if point == exclude {
			continue
		} else {
			other = append(other, point)
		}
	}
	other = append(other, include)
	return &Line{other}
}

type Cluster struct {
	base *Point
	m    map[string]*Line
}

func (this *Cluster) maxLength(repeatedMap *RepeatedMap) int {
	max := 0
	for _, line := range this.m {
		lineLength := line.length(repeatedMap)
		if lineLength > max {
			max = lineLength
		}
	}
	baseRepeatTimes, _ := repeatedMap.m[this.base.buildKey()]
	return max + baseRepeatTimes
}

func (this *Cluster) cut(repeatedMap *RepeatedMap) {
	slopes := make([]string, 0)
	for slope, line := range this.m {
		if line.length(repeatedMap) <= 1 {
			slopes = append(slopes, slope)
		}
	}
	for _, slope := range slopes {
		delete(this.m, slope)
	}
}

func (this *Cluster) findLine(slope string) (line *Line, ok bool) {
	line, ok = this.m[slope]
	return
}

func (this *Cluster) putLine(line *Line) {
	slope := this.base.slope(line.list[0])
	this.m[slope] = line
}

func (this *Cluster) putPoint(point *Point) {
	slope := this.base.slope(point)
	_, ok := this.m[slope]
	if !ok {
		this.m[slope] = &Line{[]*Point{point}}
	} else {
		this.m[slope].append(point)
	}
}

type PointSet struct {
	m map[string]*Point
}

func (this *PointSet) putLine(line *Line) bool {
	for _, point := range line.list {
		this.putPoint(point)
	}
	return true
}

func (this *PointSet) putPoint(point *Point) bool {
	_, ok := this.m[point.buildKey()]
	if ok {
		return false
	} else {
		this.m[point.buildKey()] = point
		return true
	}
}

func (this *PointSet) isExisting(point *Point) bool {
	key := point.buildKey()
	_, ok := this.m[key]
	return ok
}

type PointMap struct {
	m map[*Point]*Cluster
}

func (this *PointMap) put(point *Point) bool {
	_, ok := this.m[point]
	if ok {
		return true
	} else {
		this.m[point] = &Cluster{point, make(map[string]*Line)}
		return false
	}
}

func (this *PointMap) find(point *Point) (cluster *Cluster, ok bool) {
	cluster, ok = this.m[point]
	return
}

type RepeatedMap struct {
	m map[string]int
}

func (this *RepeatedMap) put(point *Point) (isFirstTime bool) {
	key := point.buildKey()
	_, ok := this.m[key]
	if ok {
		this.m[key]++
		isFirstTime = false
	} else {
		this.m[point.buildKey()] = 1
		isFirstTime = true
	}
	return
}

func (this *RepeatedMap) getMaxRepeated() (key string, times int) {
	times = 0
	for k, v := range this.m {
		if v > times {
			key = k
			times = v
		}
	}
	return
}

func build(points [][]int) (result []*Point, repeatedMap *RepeatedMap) {
	result = make([]*Point, 0)
	repeatedMap = &RepeatedMap{make(map[string]int)}
	for _, p := range points {
		point := &Point{p[0], p[1]}
		isFirstTime := repeatedMap.put(point)
		if isFirstTime {
			result = append(result, point)
		}
	}
	return
}

func maxPoints(arr [][]int) int {
	points, repeatedMap := build(arr)
	if len(arr) == 0 {
		return 0
	} else if len(arr) <= 2 {
		return len(arr)
	}
	pointMap := &PointMap{make(map[*Point]*Cluster)}
	max := 2
	_, maxRepeatedTimes := repeatedMap.getMaxRepeated()
	if maxRepeatedTimes > max {
		max = maxRepeatedTimes
	}
	// threshold := len(arr)
	threshold := len(points)/2 + maxRepeatedTimes
	if len(points) == 3 {
		threshold = 3
	}
	for i, cur := range points {
		pointMap.put(cur)
		curCluster := pointMap.m[cur]
		pointSet := &PointSet{make(map[string]*Point)}
		for j, point := range points {
			if i == j {
				continue
			} else if i > j {
				cluster, ok := pointMap.find(point)
				slope := point.slope(cur)
				line, ok := cluster.findLine(slope)
				if ok {
					pointSet.putLine(line)
					pointSet.putPoint(point)
					copiedLine := line.copy(point, cur)
					curCluster.putLine(copiedLine)
				}
			} else {
				if pointSet.isExisting(point) {
					continue
				}
				curCluster.putPoint(point)
			}
		}
		// curCluster.cut(repeatedMap)
		maxLineLen := curCluster.maxLength(repeatedMap)
		if maxLineLen > max {
			max = maxLineLen
		}
		if max >= threshold {
			return max
		}
	}
	return max
}

func test(points [][]int, expected int) {
	actual := maxPoints(points)
	if actual != expected {
		fmt.Printf("%v should be %d, actual %d\n", points, expected, actual)
	}
}

func main() {
	test([][]int{[]int{1, 1}}, 1)
	test([][]int{[]int{1, 1}, []int{2, 2}}, 2)
	test([][]int{[]int{1, 1}, []int{3, 2}, []int{5, 3}, []int{4, 1}, []int{2, 3}, []int{1, 4}}, 4)
	test([][]int{[]int{1, 1}, []int{2, 2}, []int{3, 3}}, 3)
	test([][]int{[]int{1, 1}, []int{2, 2}, []int{0, 3}}, 2)
	test([][]int{[]int{0, 0}, []int{1, 1}, []int{0, 0}}, 3)
	test([][]int{[]int{1, 1}, []int{1, 1}, []int{2, 2}, []int{2, 2}}, 4)
	test([][]int{[]int{-4, 1}, []int{-7, 7}, []int{-1, 5}, []int{9, -25}}, 3)
	test([][]int{[]int{0, 0}, []int{94911151, 94911150}, []int{94911152, 94911151}}, 2)
}
