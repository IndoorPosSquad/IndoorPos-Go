package positioning

import (
	"testing"
	"../mat"
)

func Test_Get_probable_pos(t *testing.T) {
	A := mat.Zeros(3, 3)
	A[0] = []float64{116.359820, 39.960088, 70}
	A[1] = []float64{116.360624, 39.960109, 55}
	A[2] = []float64{116.359772, 39.961054, 60}
	data := mat.Zeros(1, 3)
	should_be := mat.Zeros(1, 3)
	should_be[0] = []float64{116.36002568447782, 39.96035462317992, 100.2259350147533}
	data[0] = []float64{45.774405594486325, 73.42235545048611, 89.65644580836928}
	
	prob_pos, _ := Get_probable_pos(data, A, 0)
	
	if mat.Identity(prob_pos, should_be) {
		t.Log("测试通过了")
	} else {
		t.Error(prob_pos,should_be)
    }
	
}

func Test_Pos_solu(t *testing.T) {
	Satellites := mat.Zeros(3, 3)
	Satellites[0] = []float64{116.359820, 39.960088, 70}
	Satellites[1] = []float64{116.360624, 39.960109, 55}
	Satellites[2] = []float64{116.359772, 39.961054, 60}
	data := mat.Zeros(1, 3)
	data[0] = []float64{45.774405594486325, 73.42235545048611, 89.65644580836928}
	solution := Pos_solu(data, Satellites)
	should_be := mat.Zeros(1, 3)
	should_be[0] = []float64{116.36002499999994, 39.96035699999994, 100.00000001419397}
	if mat.Identity(solution, should_be) {
		t.Log("测试通过了")
	} else {
		t.Error(solution,should_be)
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