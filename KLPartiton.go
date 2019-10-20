package klpartitinlin

import (
	"errors"
	"math/rand"

	boolmatrixlib "github.com/Rakiiii/goBoolMatrix"
	graphlib "github.com/Rakiiii/goGraph"
)

type KLSolution struct {
	Solution boolmatrixlib.BoolMatrix

	Graph *graphlib.Graph

	AEdgesDifferens []IntPair
	BEdgesDifferens []IntPair
}

func (k *KLSolution) CountParameter() (int64, error) {
	result := int64(0)
	amV := k.Graph.AmountOfVertex()

	if amV != k.Solution.Heigh() {
		return result, errors.New("BoolMatrix heigh is not equls to amount of graphs vertexes")
	}
	w := k.Solution.Width()

	for j := 0; j < w; j++ {
		for i := 0; i < amV; i++ {
			if k.Solution.GetBool(i, j) {
				edges := k.Graph.GetEdges(i)
				for _, edge := range edges {
					if k.Solution.GetBool(edge, j) {
						result++
					}
				}
			}
		}
	}
	return result, nil
}

func (k *KLSolution) Init(graph *graphlib.Graph) {
	k.Solution.Init(2, graph.AmountOfVertex())
	k.Graph = graph
	for i := 0; i < graph.AmountOfVertex()/2; i++ {
		k.Solution.SetBool(rand.Intn(graph.AmountOfVertex()), 0, true)
	}
	for i := 0; i < graph.AmountOfVertex(); i++ {
		if !k.Solution.GetBool(i, 0) {
			k.Solution.SetBool(i, 1, true)
		}
	}

	k.InitDiff()

}

func (k *KLSolution) InitDiff() {
	k.AEdgesDifferens = make([]IntPair, k.Graph.AmountOfVertex()/2)
	k.BEdgesDifferens = make([]IntPair, k.Graph.AmountOfVertex()-k.Graph.AmountOfVertex()/2)
	aPoint := 0
	bPoint := 0
	for i := 0; i < k.Graph.AmountOfVertex(); i++ {
		if k.Solution.GetBool(i, 0) {
			tmp := 0
			for _, edge := range k.Graph.GetEdges(i) {
				if k.Solution.GetBool(edge, 0) {
					tmp--
				} else {
					tmp++
				}
			}
			k.AEdgesDifferens[aPoint] = IntPair{Number: i, Diff: tmp}
			aPoint++
		} else {
			tmp := 0
			for _, edge := range k.Graph.GetEdges(i) {
				if k.Solution.GetBool(edge, 1) {
					tmp--
				} else {
					tmp++
				}
			}
			k.BEdgesDifferens[bPoint] = IntPair{Number: i, Diff: tmp}
			bPoint++
		}
	}

	QuickSortIntPair(&k.AEdgesDifferens, 0, len(k.AEdgesDifferens))
	QuickSortIntPair(&k.BEdgesDifferens, 0, len(k.BEdgesDifferens))
}

func (k *KLSolution) FindBestPair() (int, int, int) {
	it := len(k.BEdgesDifferens) - 2
	for k.BEdgesDifferens[it].Diff >= k.BEdgesDifferens[len(k.BEdgesDifferens)-1].Diff-1 {
		flag := true
		for _, edge := range k.Graph.GetEdges(k.BEdgesDifferens[it].Number) {
			if edge == k.AEdgesDifferens[len(k.AEdgesDifferens)-1].Number {
				flag = false
			}
		}
		if flag {
			return k.AEdgesDifferens[len(k.AEdgesDifferens)-1].Number,
				k.AEdgesDifferens[it].Number,
				k.BEdgesDifferens[it].Diff + k.AEdgesDifferens[len(k.AEdgesDifferens)-1].Diff
		}
		it--
	}
	return k.AEdgesDifferens[len(k.AEdgesDifferens)-1].Number,
		k.BEdgesDifferens[len(k.BEdgesDifferens)-1].Number,
		k.BEdgesDifferens[len(k.BEdgesDifferens)-1].Diff + k.AEdgesDifferens[len(k.AEdgesDifferens)-1].Diff - 2
}

func (k *KLSolution) RemoveVertexFromDifferrence(av, bv int) error {
	avNum := -1
	for i, n := range k.AEdgesDifferens {
		if n.Number == av {
			avNum = i
		}
	}
	if avNum != -1 {
		k.AEdgesDifferens = append(k.AEdgesDifferens[:avNum], k.AEdgesDifferens[avNum+1:]...)
	} else {
		return errors.New("No vertex with such number in first differrenc")
	}
	bvNum := -1
	for i, n := range k.BEdgesDifferens {
		if n.Number == av {
			bvNum = i
		}
	}
	if bvNum != -1 {
		k.BEdgesDifferens = append(k.BEdgesDifferens[:avNum], k.BEdgesDifferens[avNum+1:]...)
	} else {
		return errors.New("No vertex with such number in second differrenc")
	}
	return nil
}

func (k *KLSolution) DecrementDiff(av, bv int) {
	aEdges := k.Graph.GetEdges(av)
	for _, edge := range aEdges {
		dec := CheckNumber(edge, k.AEdgesDifferens)
		if dec != -1 {
			k.AEdgesDifferens[dec].Diff--
		} else {
			inc := CheckNumber(edge, k.BEdgesDifferens)
			if inc != -1 {
				k.BEdgesDifferens[inc].Diff++
			}
		}
	}
	bEdges := k.Graph.GetEdges(av)
	for _, edge := range bEdges {
		dec := CheckNumber(edge, k.AEdgesDifferens)
		if dec != -1 {
			k.AEdgesDifferens[dec].Diff++
		} else {
			inc := CheckNumber(edge, k.BEdgesDifferens)
			if inc != -1 {
				k.BEdgesDifferens[inc].Diff--
			}
		}
	}
}
