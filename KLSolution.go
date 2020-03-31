package klpartitinlin

import (
	"errors"
	"math/rand"
	"time"

	boolmatrixlib "github.com/Rakiiii/goBoolMatrix"
	graphlib "github.com/Rakiiii/goGraph"
)

//KLSolution struct for solving graph partitioning with KL algorithm
//contains BoolMatrix as solution, graph for partitioning and edges diff
type KLSolution struct {
	Solution boolmatrixlib.BoolMatrix

	Graph graphlib.IGraph

	AEdgesDifferens []IntPair
	BEdgesDifferens []IntPair
}

//CountParameter returns value of param for setted solution
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
	return int64(k.Graph.AmountOfEdges()) - result/2, nil
}

//Init initialize grpah with @grpah param, and generating random solution
func (k *KLSolution) Init(graph graphlib.IGraph) {
	k.Solution.Init(2, graph.AmountOfVertex())
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	k.Graph = graph
	for i := 0; i < graph.AmountOfVertex()/2; i++ {
		rndIndex := rnd.Intn(graph.AmountOfVertex())
		for k.Solution.GetBool(rndIndex, 0) {
			rndIndex = rnd.Intn(graph.AmountOfVertex())
		}
		k.Solution.SetBool(rndIndex, 0, true)
	}
	for i := 0; i < graph.AmountOfVertex(); i++ {
		if !k.Solution.GetBool(i, 0) {
			k.Solution.SetBool(i, 1, true)
		}
	}

	k.InitDiff()

}

//InitDiff initializing differense betewen out and in vertex
func (k *KLSolution) InitDiff() {

	k.AEdgesDifferens = make([]IntPair, 0)
	k.BEdgesDifferens = make([]IntPair, 0)
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
			k.AEdgesDifferens = append(k.AEdgesDifferens, IntPair{Number: i, Diff: tmp})
		} else {
			tmp := 0
			for _, edge := range k.Graph.GetEdges(i) {
				if k.Solution.GetBool(edge, 1) {
					tmp--
				} else {
					tmp++
				}
			}
			k.BEdgesDifferens = append(k.BEdgesDifferens, IntPair{Number: i, Diff: tmp})
		}
	}

	k.AEdgesDifferens = QuicksortIntPair(k.AEdgesDifferens)
	k.BEdgesDifferens = QuicksortIntPair(k.BEdgesDifferens)
}

//FindBestPair returns best pair for swap in setted solution
func (k *KLSolution) FindBestPair() (int, int, int) {
	if len(k.AEdgesDifferens) < 1 || len(k.BEdgesDifferens) < 1 {
		return 0, 0, -1
	}
	itA := len(k.AEdgesDifferens) - 1
	itB := len(k.BEdgesDifferens) - 1
	maxdif, va, vb := -1, -1, -1

	for k.AEdgesDifferens[itA].Diff >= k.AEdgesDifferens[len(k.AEdgesDifferens)-1].Diff-2 && itA > 0 {
		for k.BEdgesDifferens[itB].Diff >= k.BEdgesDifferens[len(k.BEdgesDifferens)-1].Diff-2 && itB > 0 {
			flag := false
			for _, edge := range k.Graph.GetEdges(k.BEdgesDifferens[itB].Number) {
				if edge == k.AEdgesDifferens[itA].Number {
					//fmt.Println("flag", , "|", k.AEdgesDifferens[itA].Number)
					flag = true
				}
			}
			if flag {
				if maxdif < k.AEdgesDifferens[itA].Diff+k.BEdgesDifferens[itB].Diff-2 {

					va, vb, maxdif = k.AEdgesDifferens[itA].Number,
						k.BEdgesDifferens[itB].Number,
						k.AEdgesDifferens[itA].Diff+k.BEdgesDifferens[itB].Diff-2
					//fmt.Println(va, "/", vb, "/", maxdif, "}")
				}
			} else {
				if maxdif < k.AEdgesDifferens[itA].Diff+k.BEdgesDifferens[itB].Diff {
					va, vb, maxdif = k.AEdgesDifferens[itA].Number,
						k.BEdgesDifferens[itB].Number,
						k.AEdgesDifferens[itA].Diff+k.BEdgesDifferens[itB].Diff
					//fmt.Println(va, "/", vb, "/", maxdif)
				}
			}

			itB--
		}
		itA--
	}
	return va, vb, maxdif
}

//RemoveVertexFromDifferrence removes vertex for enable for swapping slice
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
		if n.Number == bv {
			bvNum = i
		}
	}
	if bvNum != -1 {
		k.BEdgesDifferens = append(k.BEdgesDifferens[:bvNum], k.BEdgesDifferens[bvNum+1:]...)
	} else {
		return errors.New("No vertex with such number in second differrenc")
	}
	return nil
}

//DecrementDiff decrementing differense in out edges for vertex with number @av param and @bv param,@av param should be in AEdgesDifference and
//@bv param must be in BEdgesDifferens
func (k *KLSolution) DecrementDiff(av, bv int) {
	aEdges := k.Graph.GetEdges(av)
	for _, edge := range aEdges {
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
	bEdges := k.Graph.GetEdges(bv)
	for _, edge := range bEdges {
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
	if len(k.AEdgesDifferens) > 1 {
		MergeSort(k.AEdgesDifferens)
	}
	if len(k.BEdgesDifferens) > 1 {
		MergeSort(k.BEdgesDifferens)
	}
}

//SwapVeretx swaps vertex in solution numbers of vertex passed in @av and @bv param
func (k *KLSolution) SwapVertex(av, bv int) {
	k.Solution.SetBool(av, 0, false)
	k.Solution.SetBool(av, 1, true)
	k.Solution.SetBool(bv, 0, true)
	k.Solution.SetBool(bv, 1, false)
}
