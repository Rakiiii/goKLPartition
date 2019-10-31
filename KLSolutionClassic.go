package klpartitinlin

import (
	"log"
	"math/rand"
	"time"

	graphlib "github.com/Rakiiii/goGraph"
)

//KLSolutionClassic struct for solving graph partitioning with KL classic algorithm
//KLSolution as inclusion and redefine some methods
type KLSolutionClassic struct {
	KLSolution
}

//Init initialize grpah with @grpah param, and generating random solution
func (k *KLSolutionClassic) Init(graph *graphlib.Graph) {
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
func (k *KLSolutionClassic) InitDiff() {
	log.Println("initDiff")
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
}

//DecrementDiff decrementing differense in out edges for vertex with number @av param and @bv param,@av param should be in AEdgesDifference and
//@bv param must be in BEdgesDifferens
func (k *KLSolutionClassic) DecrementDiff(av, bv int) {
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
}

//FindBestPair returns best pair for swap in setted solution
func (k *KLSolutionClassic) FindBestPair() (int, int, int) {
	maxdif, av, bv := -1, -1, -1
	subdif := 0
	for _, aelem := range k.AEdgesDifferens {
		for _, belem := range k.BEdgesDifferens {
			subdif = aelem.Diff + belem.Diff
			if CheckInclusion(aelem.Number, k.Graph.GetEdges(belem.Number)) {
				subdif -= 2
			}
			if subdif > maxdif {
				maxdif, av, bv = subdif, aelem.Number, belem.Number
			}
		}
	}

	return av, bv, maxdif
}
