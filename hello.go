package main

import "fmt"

func unique(arr []int) []int {
	result := []int{}
	if len(arr) == 0 {
		return []int{}
	}
	result = append(result, arr[0])
	for i := 1; i < len(arr); i += 1 {
		if result[len(result)-1] != arr[i] {
			result = append(result, arr[i])
		}
	}
	return result
}

// func notContains(arr []int, ele int) bool {
// 	result := true
// 	for _, v := range arr {
// 		if ele != v {
// 			result = false
// 		} else {
// 			result = true
// 		}
// 	}
// 	return result
// }

func contains(arr []int, value int) bool {
	result := false
	for _, v := range arr {
		if value == v {
			result = true
			break
		}
	}
	return result
}

func notContains(arr []int, ele int) bool {
	result := true
	for _, v := range arr {
		if ele == v {
			result = false
			break
		}
	}
	return result
}

func unsortedUnique(arr []int) []int {
	result := []int{}
	if len(arr) == 0 {
		return []int{}
	}
	result = append(result, arr[0])
	for i := 1; i < len(arr); i += 1 {
		exists := false
		for j := 0; j < len(result); j += 1 {
			if arr[i] == result[j] {
				exists = true
			}
		}
		if !exists {
			result = append(result, arr[i])
		}
	}
	return result
}

func intersection(ar1 []int, ar2 []int) []int {
	result := []int{}
	for i := 0; i < len(ar1); i += 1 {
		for j := 0; j < len(ar2); j += 1 {
			if ar1[i] == ar2[j] {
				result = append(result, ar1[i])
			}
		}
	}
	return result
}

func difference(ar1 []int, ar2 []int) []int {
	result := []int{}
	for i := 0; i < len(ar1); i += 1 {
		j := 0
		for ; j < len(ar2); j += 1 {
			if ar1[i] == ar2[j] {
				break
			}
		}
		if j == len(ar2) {
			result = append(result, ar1[i])
		}
	}
	return result
}

func union(ar1 []int, ar2 []int) []int {
	result := []int{}
	result = append(result, ar1...)
	for i := 0; i < len(ar2); i += 1 {
		exists := false
		for j := 0; j < len(result); j += 1 {
			if ar2[i] == result[j] {
				exists = true
			}
		}
		if !exists {
			result = append(result, ar2[i])
		}
	}
	return result
}

func sequenceSum(arr []int) []int {
	result := []int{}
	for i := 0; i < len(arr)-1; i += 1 {
		sum := arr[i] + arr[i+1]
		if sum >= 10 {
			result = append(result, sum)
		}
	}
	return result
}

func sliceOfSum(arr []int) [][2]int {
	result := [][2]int{}
	for i := 0; i < len(arr)-1; i += 1 {
		sum := arr[i] + arr[i+1]
		if sum >= 10 {
			newArr := [2]int{arr[i], arr[i+1]}
			result = append(result, newArr)
		}
	}
	return result
}

func factors(num int) []int {
	result := []int{}
	for i := 1; i <= num; i += 1 {
		if num%i == 0 {
			result = append(result, i)
		}
	}
	return result
}

func perfectNumber(num int) bool {
	sum := 0
	for i := 1; i < num; i += 1 {
		if num%i == 0 {
			sum += i
		}
	}
	return sum == num
}

func letterCount(arr []string) map[string]int {
	counts := make(map[string]int)
	for _, v := range arr {
		counts[v] += 1
	}
	return counts
}

func grouping(arr []int) [][]int {
	result := [][]int{}
	newMap := make(map[int][]int)
	for _, v := range arr {
		newMap[v] = append(newMap[v], v)
	}
	for _, value := range newMap {
		result = append(result, value)
	}
	return result
}

func divideIntoSlices(arr []int) [3][]int {
	result := [3][]int{}
	first := arr[0]
	for i := 0; i < len(arr); i += 1 {
		if first < arr[i] {
			result[0] = append(result[0], arr[i])
		} else if first > arr[i] {
			result[2] = append(result[2], arr[i])
		} else {
			result[1] = append(result[1], arr[i])
		}
	}
	return result
}

func sorted(ar1 []int, ar2 []int) []int {
	result := []int{}
	for _, v := range ar1 {
		for _, e := range ar2 {
			if v < e {
				if result[len(result)-1] != v {
					result = append(result, v)
				} else {
					continue
				}
			} else {
				result = append(result, e)
			}
		}
	}
	return result
}

func main() {
	x := []int{1, 1, 2, 2, 3, 3, 3, 4}
	y := []int{2, 2, 3}
	z := []int{1, 2, 1, 2, 4}
	j := []int{1, 2, 5, 4, 2, 1, 3, 2, 3}
	k := []int{1, 1}
	c := []int{1, 2}
	d := []int{}
	seq := []int{5, 4, 7, 3, 2, 8, 4, 1, 2, 9}
	data := []string{"a", "c", "e", "a", "e"}
	res := []int{1, 2, 1, 1, 3, 2, 2, 3, 7, 1, 9}
	src := []int{4, 7, 3, 4, 8, 3, 1, 10}
	s1 := []int{1, 2, 4, 7}
	s2 := []int{0, 3, 5, 6, 8, 9}
	fmt.Println(s2[len(s2)-1], "last")
	fmt.Println(divideIntoSlices(src), "divideIntoSlices")
	fmt.Println(grouping(res), "grouping")
	fmt.Println(letterCount(data), "letterCount")
	fmt.Println(factors(30), "factors")
	fmt.Println(perfectNumber(9), "Perfect Number")
	fmt.Println(sequenceSum(seq), "sequence")
	fmt.Println(sliceOfSum(seq), "sliceOfSum")
	d = append(d, c...)
	fmt.Println(d, "d")
	ar1 := []int{0, 1, 3, 4, 7, 8}
	ar2 := []int{9, 10, 0, 8, 3, 40}
	ar3 := []int{7, 8, 0, 3, 5}
	a1 := []int{1, 2, 4, 7, 9, 3}
	a2 := []int{5, 3, 6, 2, 8}

	// res := []int{}
	// fmt.Println(res[len(res)-1])
	fmt.Println(union(a2, a1), "union")
	fmt.Println(intersection(ar1, ar2), "intersection")
	fmt.Println(difference(ar1, ar3), "difference")
	fmt.Println(unique(x))
	fmt.Println(notContains(y, 7))
	fmt.Println(unsortedUnique(z))
	fmt.Println(unsortedUnique(j))
	fmt.Println(unsortedUnique(k))
	fmt.Println(sorted(s1, s2), "sorted")
	fmt.Println("Hello, World!")
}
