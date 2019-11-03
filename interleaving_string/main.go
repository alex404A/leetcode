package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	cache := make(map[string]bool)
	return check(s1, s2, s3, &cache)
}

func check(s1 string, s2 string, s3 string, cachePtr *map[string]bool) bool {
	cache := *cachePtr
	key := assemble(s1, s2, s3)
	isExist, ok := cache[key]
	if ok {
		return isExist
	}
	if len(s1) == 0 {
		cache[key] = s2 == s3
		return cache[key]
	}
	if len(s2) == 0 {
		cache[key] = s1 == s3
		return cache[key]
	}
	isPath1Exist := false
	isPath2Exist := false
	if s1[0] == s3[0] {
		isPath1Exist = check(s1[1:], s2, s3[1:], cachePtr)
		if isPath1Exist {
			cache[key] = isPath1Exist
			return cache[key]
		}
	}
	if s2[0] == s3[0] {
		isPath2Exist = check(s1, s2[1:], s3[1:], cachePtr)
		if isPath2Exist {
			cache[key] = isPath2Exist
			return cache[key]
		}
	}
	cache[key] = false
	return cache[key]
}

func assemble(s1 string, s2 string, s3 string) string {
	return s1 + ":" + s2 + ":" + s3
}

func test(s1 string, s2 string, s3 string, expected bool) {
	actual := isInterleave(s1, s2, s3)
	if actual != expected {
		if expected {
			fmt.Printf("%s and %s fail to interleave %s\n", s1, s2, s3)
		} else {
			fmt.Printf("%s and %s success to interleave %s\n", s1, s2, s3)
		}
	}
}

func main() {
	test("aabcc", "dbbca", "aadbbcbcac", true)
	test("aabcc", "dbbca", "aadbbbaccc", false)
}
