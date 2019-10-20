package klpartitinlin

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

func QuickSortIntPair(arr *[]IntPair, _st, _en int) {
	st := _st
	en := _en - 1

	p := (*arr)[_en>>1]
	for ok := true; ok; ok = (st <= en) {
		for (*arr)[st].Diff < p.Diff {
			st++
		}
		for (*arr)[st].Diff > p.Diff {
			en--
		}

		if st <= en {
			temp := (*arr)[st]
			(*arr)[st] = (*arr)[en]
			(*arr)[en] = temp
			st++
			en--
		}
	}

	if en > 0 {
		QuickSortIntPair(arr, 0, en)
	}
	if _en > st {
		QuickSortIntPair(arr, st, _en-st)
	}

}

func MergeSortIntPair(arr *[]IntPair, _st, _en int) {
	if _st < _en {
		split := (_st + _en) / 2
		MergeSortIntPair(arr, _st, split)
		MergeSortIntPair(arr, split, _en)
		Merge(arr, _st, split, _en)
	}

}

func Merge(arr *[]IntPair, _st, _sp, _en int) {
	posSt := _st
	posEn := _sp + 1

	posRes := 0

	res := make([]IntPair, _en-_st+1)

	for posSt <= _sp && posEn <= _en {
		if (*arr)[posSt].Diff < (*arr)[posEn].Diff {
			res[posRes] = (*arr)[posSt]
			posSt++
		} else {
			res[posRes] = (*arr)[posEn]
			posEn++
		}
		posRes++
	}

	for posEn <= _en {
		res[posRes] = (*arr)[posEn]
		posRes++
		posEn++
	}
	for posSt <= _sp {
		res[posRes] = (*arr)[posSt]
		posSt++
		posRes++
	}

	for posRes = 0; posRes < _en-_st+1; posRes++ {
		(*arr)[_st+posRes] = res[posRes]
	}
}
