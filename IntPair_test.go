package klpartitinlin

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"testing"
	"time"

	graphlib "github.com/Rakiiii/goGraph"
)

func TestKLAlgorithm(t *testing.T) {
	var parser = new(graphlib.Parser)
	var g, err = parser.ParseUnweightedUndirectedGraphFromFile("test_gr3")
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("graph parsed")

	maxTime, _ := time.ParseDuration("120ms")

	result := Result{Matrix: nil, Value: -1}

	var prevValue int64 = math.MaxInt64
	startTime := time.Now()
	endTime := time.Now()

	firstItTime, _ := time.ParseDuration("0ms")
	for endTime.Sub(startTime).Milliseconds() < maxTime.Milliseconds() {

		log.Println("start partition")
		result, err = KLPartitionigAlgorithm(g, result.Matrix)

		if err != nil {
			log.Println(err)
			return
		}

		endTime = time.Now()
		if firstItTime.Milliseconds() == 0 {
			firstItTime = endTime.Sub(startTime)
		}

		if err != nil {
			log.Println(err)
			return
		}

		if prevValue <= result.Value {
			writeTime(firstItTime)
			saveResult(&result)
			return
		} else {
			prevValue = result.Value
		}

	}

	writeTime(firstItTime)
	saveResult(&result)
}

func TestQuicksortIntPair(t *testing.T) {
	s := make([]IntPair, 10)
	for i, _ := range s {
		s[9-i].Diff = i
		s[i].Number = i
	}
	PrintIntPairSlice(s)
}

func writeTime(time time.Duration) {
	timeFile, err := os.Create("time")
	defer timeFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		timeFile.WriteString(strconv.FormatInt(time.Milliseconds(), 10) + "ms")
	}
}

func saveResult(result *Result) {
	f, err := os.Create("result_")
	if err != nil {
		fmt.Println(err)
		fmt.Println(result.Value)
		for i := 0; i < result.Matrix.Heigh(); i++ {
			for j := 0; j < result.Matrix.Width(); j++ {
				fmt.Print(result.Matrix.GetBool(i, j))
			}
			fmt.Println()
		}
		return
	}
	defer f.Close()

	f.WriteString(strconv.FormatInt(result.Value, 10) + "\n")
	for i := 0; i < result.Matrix.Heigh(); i++ {
		subStr := ""
		for j := 0; j < result.Matrix.Width(); j++ {
			if result.Matrix.GetBool(i, j) {
				subStr = subStr + string("1 ")
			} else {
				subStr = subStr + string("0 ")
			}
		}
		subStr = subStr + "\n"
		f.WriteString(subStr)

	}

}
