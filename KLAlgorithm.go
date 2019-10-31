package klpartitinlin

import (
	"log"

	boolmatrixlib "github.com/Rakiiii/goBoolMatrix"
	graphlib "github.com/Rakiiii/goGraph"
)

//Resutl struct contains matrix of graph partition and value of parameter of partition
type Result struct {
	Matrix *boolmatrixlib.BoolMatrix
	Value  int64
}

//KLPartitionigAlgorithm algorithm of graph partition(fast realisation) grtting link to graph as param @grpah,
//link to  previosu soluiotn as @_sol param(set nil if first itteration)
func KLPartitionigAlgorithm(graph *graphlib.Graph, _sol *boolmatrixlib.BoolMatrix) (Result, error) {
	result := Result{Matrix: nil, Value: -1}
	sol := new(KLSolution)
	if _sol != nil {
		sol.Graph = graph
		sol.Solution = *_sol
		sol.InitDiff()
	} else {
		sol.Init(graph)
	}

	val, er := sol.CountParameter()
	if er != nil {
		return result, er
	}
	log.Println("start Value:", val)

	fVertex, sVertex, dif := sol.FindBestPair()

	swapVertex := make([]int, 0)
	for dif > 0 {

		swapVertex = append(swapVertex, fVertex)
		swapVertex = append(swapVertex, sVertex)

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

	return result, nil
}
