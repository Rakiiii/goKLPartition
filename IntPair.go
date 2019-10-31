package klpartitinlin

import (
	"fmt"
)

type IntPair struct {
	Number, Diff int
}

//CheckNumber checks is @arr containsa @num
func CheckNumber(num int, arr []IntPair) int {
	for i, elem := range arr {
		if elem.Number == num {
			return i
		}
	}
	return -1
}

//Print printing *@p to stdout
func (p *IntPair) Print() {
	fmt.Print(p.Number, "|", p.Diff)
}

//Println printing *@p to stdout from new line
func (p *IntPair) Println() {
	fmt.Println(p.Number, "|", p.Diff)
}

//PrintIntPairSlice prints @s to stdout
func PrintIntPairSlice(s []IntPair) {
	for _, i := range s {
		i.Println()
	}
}

//QuicksortIntPair sort @a by quick sort algorithm by diff incresment
func QuicksortIntPair(a []IntPair) []IntPair {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := len(a) >> 1
	//pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i].Diff < a[right].Diff {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuicksortIntPair(a[:left])
	QuicksortIntPair(a[left+1:])

	return a
}

//MergeSort sort @a by merge sort algorithm by diff incresment
func MergeSort(items []IntPair) []IntPair {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]IntPair, middle)
		right = make([]IntPair, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return Merge(MergeSort(left), MergeSort(right))
}

//Merge returns result of merging @left and @right
func Merge(left, right []IntPair) (result []IntPair) {
	result = make([]IntPair, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0].Diff < right[0].Diff {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
