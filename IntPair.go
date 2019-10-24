package klpartitinlin

import (
	"fmt"
)

type IntPair struct {
	Number, Diff int
}

func CheckNumber(num int, arr []IntPair) int {
	for i, elem := range arr {
		if elem.Number == num {
			return i
		}
	}
	return -1
}

func (p *IntPair) Print() {
	fmt.Print(p.Number, "|", p.Diff)
}
func (p *IntPair) Println() {
	fmt.Println(p.Number, "|", p.Diff)
}

func PrintIntPairSlice(s []IntPair) {
	for _, i := range s {
		i.Println()
	}
}

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
