package klpartitinlin

import (
	"log"
	"math/rand"
	"time"

	graphlib "github.com/Rakiiii/goGraph"
)

type KLSolutionClassic struct {
	KLSolution
}

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

func (k *KLSolutionClassic) InitDiff() {
	log.Println("initDiff")
	k.AEdgesDifferens = make([]IntPair, 0)
	k.BEdgesDifferens = make([]IntPair, 0)
	for i := 0; i < k.Graph.AmountOfVertex(); i++ {
		if k.Solution.GetBool(i, 0) {
			//fmt.Println(i, "a:")
			tmp := 0
			for _, edge := range k.Graph.GetEdges(i) {
				if k.Solution.GetBool(edge, 0) {
					tmp--
				} else {
					tmp++
				}
				//	fmt.Print(edge, "|", tmp, "	")
			}
			//fmt.Println()
			k.AEdgesDifferens = append(k.AEdgesDifferens, IntPair{Number: i, Diff: tmp})
		} else {
			//fmt.Println(i, "b:")
			tmp := 0
			for _, edge := range k.Graph.GetEdges(i) {
				if k.Solution.GetBool(edge, 1) {
					tmp--
				} else {
					tmp++
				}
				//	fmt.Print(edge, "|", tmp, "	")
			}
			//fmt.Println()
			k.BEdgesDifferens = append(k.BEdgesDifferens, IntPair{Number: i, Diff: tmp})
		}
	}
}

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
