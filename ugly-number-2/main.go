package main

func nthUglyNumber(n int) int {
	if n <= 1 {
		return 1
	}
	result := make([]int, n)
	result[0] = 1
	t2 := 0
	t3 := 0
	t5 := 0
	for i := 1; i < n; i++ {
		m := min(result[t2]*2, result[t3]*3, result[t5]*5)
		if result[t2]*2 == m {
			t2++
		}
		if result[t3]*3 == m {
			t3++
		}
		if result[t5]*5 == m {
			t5++
		}
		result[i] = m
	}
	return result[n-1]
}

func min(i, j, k int) int {
	if i <= j && i <= k {
		return i
	} else if j <= i && j <= k {
		return j
	} else {
		return k
	}
}

func main() {
	nthUglyNumber(10)
}
