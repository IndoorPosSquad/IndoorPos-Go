package positioning

import (
	"testing"
	"../mat"
	"math"
)

func Test_Pos_solu(t *testing.T) {
	Satellites := mat.Zeros(3, 3)
	Satellites[0] = []float64{116.359832, 39.960092, 50}
	Satellites[1] = []float64{116.359580, 39.960466, 50}
	Satellites[2] = []float64{116.360836, 39.960573, 50}
	data := mat.Zeros(1, 3)
	//data[0] = []float64{116.358565, 39.961126, 150.156}
	data[0] = []float64{116.387603, 40.111207, 150.156}
	p1 := math.Sqrt(math.Pow((Satellites[0][0] - data[0][0]) * 3600 * 23.69, 2) + math.Pow((Satellites[0][1] - data[0][1]) * 3600 * 30.8, 2) + math.Pow((Satellites[0][2] - data[0][2]), 2))
	p2 := math.Sqrt(math.Pow((Satellites[1][0] - data[0][0]) * 3600 * 23.69, 2) + math.Pow((Satellites[1][1] - data[0][1]) * 3600 * 30.8, 2) + math.Pow((Satellites[1][2] - data[0][2]), 2))
	p3 := math.Sqrt(math.Pow((Satellites[2][0] - data[0][0]) * 3600 * 23.69, 2) + math.Pow((Satellites[2][1] - data[0][1]) * 3600 * 30.8, 2) + math.Pow((Satellites[2][2] - data[0][2]), 2))
	data2 := mat.Zeros(1, 3)
	data2[0] = []float64{p1, p2, p3}
	guess, _ := Get_probable_pos(data2, Satellites, 0)
	solu := Pos_solu(data2, Satellites)
	
	e1 := math.Sqrt(math.Pow((solu[0][0] - data[0][0]) * 3600 * 23.69, 2) + math.Pow((solu[0][1] - data[0][1]) * 3600 * 30.8, 2) + math.Pow((solu[0][2] - data[0][2]), 2))
	e2 := math.Sqrt(math.Pow((guess[0][0] - data[0][0]) * 3600 * 23.69, 2) + math.Pow((guess[0][1] - data[0][1]) * 3600 * 30.8, 2) + math.Pow((guess[0][2] - data[0][2]), 2))
	if e2 < 0.00000001 {
		t.Log("粗算位置测试通过了,误差为：", e2, "Guess:", guess)
	} else {
		t.Error("粗算测试没有通过！！误差为：", e2, "Guess:", guess)
    }
	if e1 < 0.00001 {
		t.Log("精确解算位置测试通过了,误差为：", e1)
	} else {
		t.Error("精确解算测试没有通过！！误差为：", e1)
    }
}

func Benchmark_Pos_solu(b *testing.B){
	data := mat.Zeros(1, 3)
	data[0] = []float64{49.774405594486325, 73.42235545048611, 89.65644580836928}
	Satellites := mat.Zeros(3, 3)
	Satellites[0] = []float64{116.359820, 39.960088, 70}
	Satellites[1] = []float64{116.360624, 39.960109, 55}
	Satellites[2] = []float64{116.359772, 39.961054, 60}
	
	for i := 0; i < b.N; i++ {
		Pos_solu(data, Satellites)
	}
}