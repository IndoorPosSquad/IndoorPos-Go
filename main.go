package main

import (
	"fmt"
	"./positioning"
	"./mat"
	"math"
)

func Solu(A []float64, B [][]float64) [][]float64 {
	data := mat.Zeros(1, 3)
	data[0] = []float64{A[0], A[1], A[2]}
	Satellites := B
	solution := positioning.Pos_solu(data, Satellites)
	return solution
}
func main() {
	Satellites := mat.Zeros(3, 3)
	Satellites[0] = []float64{116.359820, 39.960088, 70}
	Satellites[1] = []float64{116.360624, 39.960109, 55}
	Satellites[2] = []float64{116.359772, 39.961054, 60}
	data := mat.Zeros(1, 3)
	data[0] = []float64{116.360025, 39.960357, 100}
	p1 := math.Sqrt(math.Pow((Satellites[0][0] - data[0][0]) * 3600 * 23.69, 2) + math.Pow((Satellites[0][1] - data[0][1]) * 3600 * 30.8, 2) + math.Pow((Satellites[0][2] - data[0][2]), 2))
	p2 := math.Sqrt(math.Pow((Satellites[1][0] - data[0][0]) * 3600 * 23.69, 2) + math.Pow((Satellites[1][1] - data[0][1]) * 3600 * 30.8, 2) + math.Pow((Satellites[1][2] - data[0][2]), 2))
	p3 := math.Sqrt(math.Pow((Satellites[2][0] - data[0][0]) * 3600 * 23.69, 2) + math.Pow((Satellites[2][1] - data[0][1]) * 3600 * 30.8, 2) + math.Pow((Satellites[2][2] - data[0][2]), 2))
	
	data2 := []float64{p1, p2, p3}
	//49.774405594486325, 73.42235545048611, 89.65644580836928}
	
	solution := Solu(data2, Satellites)
	fmt.Println(solution)
	fmt.Println(data2)
}