package main

import (
		mat "github.com/skelterjohn/go.matrix"
		rand "math/rand"
		"math"
		"fmt"
		"sort"
		// "errors"
		)
 
//Sorts a map by value size in .s property
type sortedMap struct {
	m map[int]float64
	s []int
}
 
func (sm *sortedMap) Len() int {
	return len(sm.m)
}
 
func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}
 
func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}
 
func sortMap(m map[int]float64) []int {
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]int, len(m))
	i := 0
	for key, _ := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}

type KNNClassifier struct {
	Data mat.DenseMatrix
	Name string
	Labels []string
}

func RandomArray(n int) []float64 {
	ReturnedArray := make([]float64, n)
	for i := 0; i < n; i++ {
		ReturnedArray[i] = rand.Float64()
	}
	return ReturnedArray
}

//Mints a new classifier
func (KNN *KNNClassifier) New(name string, labels []string, numbers []float64, x int, y int) {
	
	// if x != len(KNN.Labels) {
	// 	return errors.New("KNN: There must be a label for each row")
	// }

	KNN.Data = *mat.MakeDenseMatrix(numbers, x, y)
	KNN.Name = name
	KNN.Labels = labels
}

//Computes a variety of distance metrics between two vectors
//Only returns Euclidean distance at the moment
func (KNN *KNNClassifier) ComputeDistance(vector *mat.DenseMatrix) float64 {
	var sum float64

	difference, err := KNN.Data.MinusDense(vector)
	flat := difference.Array()

	if err != nil {
		fmt.Println(err)
	}

	for _, i := range flat {
		squared := math.Pow(i, 2)
		sum += squared
	}

	eucdistance := math.Sqrt(sum)
	return eucdistance
}

//Returns a classification based on a vector input
//Just need to build the max voting function
func (KNN *KNNClassifier) Predict(vector *mat.DenseMatrix, K int) ([]string, []int) {

	rows := KNN.Data.Rows()
	rownumbers := make(map[int]float64)
	labels := make([]string, K)

	for i := 0; i < rows; i++{
		eucdistance := KNN.ComputeDistance(vector)
		rownumbers[i] = eucdistance
	}

	sorted := sortMap(rownumbers)
	values := sorted[:K]

	for _, elem := range values {
		labels = append(labels, KNN.Labels[elem])
	}

	return labels[2:], values
}

//Returns a label, given an index
func (KNN *KNNClassifier) GetLabel(index int) string {
	return KNN.Labels[index]
}

func main(){
	for {
		values := RandomArray(8)
		knn := KNNClassifier{}
		random := mat.MakeDenseMatrixStacked([][]float64{{1,4,3,2}, {1,4,3,2}})
		knn.New("Testing", []string{"One label", "Another label"}, values,2,4)
		f,d := knn.Predict(random, 2)
		fmt.Println(f,d)
	}
}