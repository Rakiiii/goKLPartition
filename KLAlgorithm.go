package klpartitinlin

import (
	boolmatrixlib "github.com/Rakiiii/goBoolMatrix"
	graphlib "github.com/Rakiiii/goGraph"
)

type Result struct {
	Matrix *boolmatrixlib.BoolMatrix
	Value  int64
}

func KLPartitionigAlgorithm(graph *graphlib.Graph, _sol *boolmatrixlib.BoolMatrix) (Result, error) {
	result := Result{Matrix: nil, Value: -1}
	sol := new(KLSolution)
	sol.Init(graph)
	if _sol != nil {
		sol.Graph = graph
		sol.Solution = *_sol
		sol.InitDiff()
	} else {
		sol.Init(graph)
	}

	/*sol.Solution.Print()
	fmt.Println()
	PrintIntPairSlice(sol.AEdgesDifferens)
	PrintIntPairSlice(sol.BEdgesDifferens)
	fmt.Println()*/

	fVertex, sVertex, dif := sol.FindBestPair()

	swapVertex := make([]int, 0)
	for dif > 0 {

		swapVertex = append(swapVertex, fVertex)
		swapVertex = append(swapVertex, sVertex)

		//fmt.Println("differrence:", dif)

		if err := sol.RemoveVertexFromDifferrence(fVertex, sVertex); err != nil {
			return result, err
		}
		sol.DecrementDiff(fVertex, sVertex)
		fVertex, sVertex, dif = sol.FindBestPair()
	}

	for i := 0; i < len(swapVertex); i += 2 {
		sol.SwapVertex(swapVertex[i], swapVertex[i+1])
	}

	var err error
	result.Value, err = sol.CountParameter()
	if err != nil {
		return result, err
	}
	result.Matrix = sol.Solution.Copy()

	//fmt.Println("value:", result.Value)

	return result, nil
}
