package main

import (
	"fmt"
)

func map1() {
	fmt.Println("--------001.map1---------")
	m := map[string]string {
		"1":"cpp",
		"2":"golang",
		"3":"java",
		"4":"python",
	}
	fmt.Println(m)

}

func map2() {
	fmt.Println("--------002.map2---------")
	m2 := make(map[string]int)//m2 == empty map
	var m3 map[string]int     //m3 == nil
	fmt.Println(m2, m3)
}

func map3() {
	fmt.Println("--------003.map3---------")
	m := map[string]string {
		"1":"cpp",
		"2":"golang",
		"3":"java",
		"4":"python",
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	v1, found1 := m["1"]
	v2, found2 := m["5"]
	fmt.Println(v1, found1)
	fmt.Println(v2, found2)
}

func deleteMap() {
	m := map[string]string {
		"1":"cpp",
		"2":"golang",
		"3":"java",
		"4":"python",
	}
	delete(m, "1")
	fmt.Println(m)
}

func findLongestSubString(str string) int {
	m := make(map[byte]int)
	left := 0
	right := -1
	maxLen := 0
	bytes := []byte(str)
	for i, c := range bytes {
		index, ok := m[c]
		if ok {
			for j:=left; j<=index; j++ {
				delete(m, bytes[j])
			}
			left = index + 1
		}
		m[c] = i
		right = i
		if right-left+1 > maxLen {
			maxLen = right-left+1
		}
	}
	fmt.Println(maxLen)
	return maxLen
}

func main() {
	map1()
	map2()
	map3()
	deleteMap()
	findLongestSubString("hello")
	findLongestSubString("bbbbb")
	findLongestSubString("pwwkew")
	findLongestSubString("")
	findLongestSubString("b")
	findLongestSubString("bac")
	findLongestSubString("abba")
}
