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
	// 操场左下3个角坐标
	Satellites[0] = []float64{116.359770, 39.961050, 60}
	Satellites[1] = []float64{116.359824, 39.960080, 50}
	Satellites[2] = []float64{116.360626, 39.960109, 70}
	
	// Satellites[0] = []float64{116.359820, 39.960088, 70}
	// Satellites[1] = []float64{116.360624, 39.960109, 70}
	// Satellites[2] = []float64{116.359772, 39.961054, 70}
	data := mat.Zeros(1, 3)
	// 毛主席
	data[0] = []float64{116.356385, 39.961186, 150.156}
	// 西门
	// data[0] = []float64{116.355243, 39.961067, 150.156}
	// 央财
	// data[0] = []float64{116.342580, 39.959169, 150.156}
	// 魏公村 北理
	// data[0] = []float64{116.322928, 39.957162, 150.156}
	// 宏福
	// data[0] = []float64{116.361404, 40.097679, 150.156}
	// 京师大厦
	// data[0] = []float64{116.371444, 39.958017, 150.156}
	// 三里屯
	// data[0] = []float64{116.464398, 39.938079, 150.156}
	// ***************************************************
	// 以上均测试通过
	// 廊坊(太远了..别忘了基站还在北邮操场上..这个算法是用
	// 体积推算高度H的，这个距离上体积基本上是0了.所以高度
	// 为0..P1,P2,P3可以近似等于平面上点到3点的距离)
	// data[0] = []float64{116.730388, 39.530102, 150.156}
	// ***************************************************
	
	p1 := math.Sqrt(math.Pow((Satellites[0][0] - data[0][0]) * 3600 * 23.69, 2) + math.Pow((Satellites[0][1] - data[0][1]) * 3600 * 30.8, 2) + math.Pow((Satellites[0][2] - data[0][2]), 2))
	p2 := math.Sqrt(math.Pow((Satellites[1][0] - data[0][0]) * 3600 * 23.69, 2) + math.Pow((Satellites[1][1] - data[0][1]) * 3600 * 30.8, 2) + math.Pow((Satellites[1][2] - data[0][2]), 2))
	p3 := math.Sqrt(math.Pow((Satellites[2][0] - data[0][0]) * 3600 * 23.69, 2) + math.Pow((Satellites[2][1] - data[0][1]) * 3600 * 30.8, 2) + math.Pow((Satellites[2][2] - data[0][2]), 2))
	
	data2 := []float64{p1, p2, p3}
	data3 := mat.Zeros(1, 3)
	data3[0] = data2
	prabo, _ := positioning.Get_probable_pos(data3, Satellites, 0)
	fmt.Println(data2)
	fmt.Println(data)
	fmt.Println(prabo)
	solution := Solu(data2, Satellites)
	fmt.Println(solution)
}