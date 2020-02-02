package main

import "fmt"

type Summary struct {
	furthest int
	balance  int
}

func (summary *Summary) isValid() bool {
	return summary.furthest != -1
}

func canCompleteCircuit(gas []int, cost []int) int {
	candidates, ok := iter(gas, cost)
	if !ok {
		return -1
	}
	m := make(map[int]Summary)
	for _, index := range candidates {
		m[index] = Summary{-1, -1}
	}
	for i := len(candidates) - 1; i >= 0; i-- {
		if check(candidates[i], gas, cost, m) {
			return candidates[i]
		}
	}
	return -1
}

func iter(gas []int, cost []int) ([]int, bool) {
	candidates := make([]int, 0)
	gasSum := 0
	costSum := 0
	for i := 0; i < len(gas); i++ {
		if gas[i] >= cost[i] {
			candidates = append(candidates, i)
		}
		gasSum += gas[i]
		costSum += cost[i]
	}
	return candidates, gasSum >= costSum
}

func check(cur int, gas []int, cost []int, m map[int]Summary) bool {
	length := len(gas)
	remaining := 0
	for i := cur; i < len(gas)+cur; i++ {
		index := getIndex(i, length)
		summary, ok := m[index]
		if ok && summary.isValid() && (cur > summary.furthest || index < summary.furthest) {
			if remaining < summary.balance {
				return false
			} else {
				i = getNext(index, summary.furthest, length)
			}
		} else {
			if remaining+gas[index]-cost[index] < 0 {
				summary := m[cur]
				summary.furthest = index
				summary.balance = cost[index] - remaining - gas[index]
				return false
			} else {
				remaining = remaining + gas[index] - cost[index]
			}
		}
	}
	return true
}

func getIndex(seq int, length int) int {
	if seq < length {
		return seq
	} else {
		return seq % length
	}
}

func getNext(index int, furthest int, length int) int {
	if index < furthest {
		return furthest
	} else {
		return length + furthest
	}
}

func main() {
	gas := []int{1, 2, 3, 4, 5, 6, 7}
	cost := []int{3, 4, 5, 1, 2, 6, 7}
	result := canCompleteCircuit(gas, cost)
	fmt.Println(result)
}
