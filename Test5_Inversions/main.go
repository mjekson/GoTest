package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type array struct {
	a []int
}

func readFile(fl string) ([]int, error) {

	b, err := os.ReadFile(fl)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	nums := make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil

}

func main() {
	nums, err := readFile("arrays.txt")
	if err != nil {
		log.Fatalf("ERROR:%v", err)
	}
	arr := array{nums}
	fmt.Println("-----Counting Inversions using Brute Force-----")
	inversions := inversionsBruteForce(&arr)
	fmt.Println("No.of Inversions =", inversions)
	fmt.Println("-----Counting Inversions using Recursion (Divide & Conquer)-----")
	sortedArr, inversions := recInversionCount(&arr)
	fmt.Println("sortedArr =", sortedArr, "and No.of Inversions =", inversions)
}

func recInversionCount(v *array) (*array, int) {
	if len(v.a) == 1 {
		return v, 0
	}
	mid := (len(v.a)) / 2
	leftHalf, leftInv := recInversionCount(&array{v.a[:mid]})
	rightHalf, rightInv := recInversionCount(&array{v.a[mid:]})
	sortedArr := []int{}
	splitInv := 0
	i, j := 0, 0
	for i < len(leftHalf.a) && j < len(rightHalf.a) {
		if leftHalf.a[i] <= rightHalf.a[j] {
			sortedArr = append(sortedArr, leftHalf.a[i])
			i++
		} else {
			sortedArr = append(sortedArr, rightHalf.a[j])
			splitInv += len(leftHalf.a) - i
			j++
		}
	}
	if i < len(leftHalf.a) {
		sortedArr = append(sortedArr, leftHalf.a[i:]...)
	}
	for j < len(rightHalf.a) {
		sortedArr = append(sortedArr, rightHalf.a[j])
		j++
	}
	return &array{sortedArr}, leftInv + rightInv + splitInv
}

/*
* Brute Force Approach
* Algorithm Complexity = O(n^2)
 */
func inversionsBruteForce(v *array) int {
	var cnt int
	for i := 0; i < len(v.a)-1; i++ {
		for j := i; j < len(v.a); j++ {
			if v.a[i] > v.a[j] {
				cnt += 1
			}
		}
	}
	return cnt
}
