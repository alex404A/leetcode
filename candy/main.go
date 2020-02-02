package main

import "fmt"

type Item struct {
	rating int
	candy  int
}

func candy(ratings []int) int {
	ratingIntervals := split(ratings)
	candyIntervals := make([][]*Item, 0)
	for _, interval := range ratingIntervals {
		candyIntervals = append(candyIntervals, distribute(interval))
	}
	candyIntervals = merge(candyIntervals)
	return sum(candyIntervals)
}

func split(ratings []int) [][]int {
	intervals := make([][]int, 0)
	interval := make([]int, 0)
	trend := "both"
	for _, rating := range ratings {
		if len(interval) == 0 {
			interval = append(interval, rating)
		} else {
			if trend == "both" {
				interval = append(interval, rating)
				if rating > interval[len(interval)-2] {
					trend = "incr"
				}
				if rating < interval[len(interval)-2] {
					trend = "decr"
				}
			} else if trend == "incr" {
				if rating >= interval[len(interval)-1] {
					interval = append(interval, rating)
				} else {
					intervals = append(intervals, interval)
					trend = "both"
					interval = []int{rating}
				}
			} else if trend == "decr" {
				if rating <= interval[len(interval)-1] {
					interval = append(interval, rating)
				} else {
					intervals = append(intervals, interval)
					trend = "both"
					interval = []int{rating}
				}
			}
		}
	}
	if len(interval) > 0 {
		intervals = append(intervals, interval)
	}
	return intervals
}

func distribute(ratings []int) []*Item {
	if len(ratings) == 1 {
		return []*Item{&Item{ratings[0], 1}}
	} else {
		candies := make([]*Item, 0)
		if ratings[len(ratings)-1] > ratings[0] {
			candies = append(candies, &Item{ratings[0], 1})
			for i := 1; i < len(ratings); i++ {
				rating := ratings[i]
				if rating == candies[len(candies)-1].rating {
					candies = append(candies, &Item{rating, 1})
				} else {
					candies = append(candies, &Item{rating, candies[len(candies)-1].candy + 1})
				}
			}
			return candies
		} else {
			candies = append(candies, &Item{ratings[len(ratings)-1], 1})
			for i := len(ratings) - 2; i >= 0; i-- {
				rating := ratings[i]
				if rating == candies[len(candies)-1].rating {
					candies = append(candies, &Item{rating, 1})
				} else {
					candies = append(candies, &Item{rating, candies[len(candies)-1].candy + 1})
				}
			}
			reversedCandies := make([]*Item, len(candies))
			for i := 0; i < len(candies); i++ {
				reversedCandies[len(candies)-1-i] = candies[i]
			}
			return reversedCandies
		}
	}
}

func merge(intervals [][]*Item) [][]*Item {
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		last := intervals[i-1]
		// 前一个是incr，当前的是decr
		if interval[0].rating < last[len(last)-1].rating {
			if interval[0].candy >= last[len(last)-1].candy {
				last[len(last)-1].candy = interval[0].candy + 1
			}
		} else {
			// 前一个是decr，当前的是incr
			if interval[0].candy <= last[len(last)-1].candy {
				addIncr(interval)
			}
		}
	}
	return intervals
}

func addIncr(interval []*Item) {
	interval[0].candy++
	for i := 1; i < len(interval); i++ {
		if interval[i].rating == interval[i-1].rating {
			return
		} else {
			interval[i].candy++
		}
	}
}

func sum(intervals [][]*Item) int {
	cnt := 0
	for _, interval := range intervals {
		for _, item := range interval {
			cnt += item.candy
		}
	}
	return cnt
}

func test(ratings []int, expected int) {
	actual := candy(ratings)
	if actual != expected {
		fmt.Printf("%v actual %d, expected %d", ratings, actual, expected)
	}
}

func main() {
	// test([]int{1, 0, 2}, 5)
	// test([]int{1, 2, 2}, 4)
	// test([]int{2, 2, 1}, 4)
	// test([]int{2, 2, 1, 1, 2, 2}, 8)
	test([]int{2, 3, 2, 1, 1, 2, 3, 4, 5}, 22)
	test([]int{2, 3, 2, 1, 1, 2, 3, 3, 5}, 16)
	test([]int{2, 3, 2, 1, 1}, 8)
}
