package positioning

import (
	"../mat"
	"math"
	"log"
)

var (
	Trans, _Trans, Satellites_pos mat.Matrix
)
const (
	NUMOFPL = 3
)

func Get_probable_pos(Distance, _Satellites_pos mat.Matrix, mode int) (POS, Satellites_xyz mat.Matrix) {

	Satellites_pos = _Satellites_pos
	
	_A := mat.CopyMatrix(Satellites_pos)
	h_2 := _A[1][2] - _A[0][2]
	_A[1][2] = _A[0][2]
	h_3 := _A[2][2] - _A[0][2]
	_A[2][2] = _A[0][2]
	
	p1 := Distance[0][0]
	p2 := Distance[0][1]
	p3 := Distance[0][2]
	
	_AB := math.Sqrt(math.Pow((_A[1][0] - _A[0][0]) * 3600 * 23.69, 2) + math.Pow((_A[1][1] - _A[0][1]) * 3600 * 30.8, 2) + math.Pow((_A[1][2] - _A[0][2]), 2))
	_AC := math.Sqrt(math.Pow((_A[2][0] - _A[0][0]) * 3600 * 23.69, 2) + math.Pow((_A[2][1] - _A[0][1]) * 3600 * 30.8, 2) + math.Pow((_A[2][2] - _A[0][2]), 2))
	_BC := math.Sqrt(math.Pow((_A[2][0] - _A[1][0]) * 3600 * 23.69, 2) + math.Pow((_A[2][1] - _A[1][1]) * 3600 * 30.8, 2) + math.Pow((_A[2][2] - _A[1][2]), 2))
	_COS_A := (math.Pow(_AC, 2) + math.Pow(_AB, 2) - math.Pow(_BC, 2)) / (2 * _AC * _AB)
	
	_Trans = mat.Zeros(3, 3)
	// X
	_Trans[0][0] = (_A[1][0] - _A[0][0]) / _AB
	_Trans[0][1] = (_A[1][1] - _A[0][1]) / _AB
	_Trans[0][2] = (_A[1][2] - _A[0][2]) / _AB
	// Y
	_d := _AC * _COS_A
	_H := math.Sqrt(_AC * _AC - _d * _d)
	var D [3]float64
	D[0] = _A[0][0] + _d * _Trans[0][0]
	D[1] = _A[0][1] + _d * _Trans[0][1]
	D[2] = _A[0][2] + _d * _Trans[0][2]
	_Trans[1][0] = (_A[2][0] - D[0]) / _H
	_Trans[1][1] = (_A[2][1] - D[1]) / _H
	_Trans[1][2] = (_A[2][2] - D[2]) / _H
	// Z
	_Trans[2][0] = 0
	_Trans[2][1] = 0
	_Trans[2][2] = 1

	Satellites_xyz = mat.Zeros(3, 3)
	Satellites_xyz[1][0] = _AB
	Satellites_xyz[1][2] = h_2
	
	Satellites_xyz[2][0] = _d
	Satellites_xyz[2][1] = _H
	Satellites_xyz[2][2] = h_3
	
	AB := math.Sqrt(math.Pow((Satellites_pos[1][0] - Satellites_pos[0][0]) * 3600 * 23.69, 2) + math.Pow((Satellites_pos[1][1] - Satellites_pos[0][1]) * 3600 * 30.8, 2) + math.Pow((Satellites_pos[1][2] - Satellites_pos[0][2]), 2))
	AC := math.Sqrt(math.Pow((Satellites_pos[2][0] - Satellites_pos[0][0]) * 3600 * 23.69, 2) + math.Pow((Satellites_pos[2][1] - Satellites_pos[0][1]) * 3600 * 30.8, 2) + math.Pow((Satellites_pos[2][2] - Satellites_pos[0][2]), 2))
	BC := math.Sqrt(math.Pow((Satellites_pos[2][0] - Satellites_pos[1][0]) * 3600 * 23.69, 2) + math.Pow((Satellites_pos[2][1] - Satellites_pos[1][1]) * 3600 * 30.8, 2) + math.Pow((Satellites_pos[2][2] - Satellites_pos[1][2]), 2))
	COS_A := (math.Pow(AC, 2) + math.Pow(AB, 2) - math.Pow(BC, 2)) / (2 * AC * AB)
	d := AC * COS_A
	H := math.Sqrt(AC * AC - d * d)
	
	Trans = mat.Zeros(3,3)

	Trans[0][0] = _AB / AB
	Trans[0][2] = h_2 / AB

	Trans[1][0] = (_d - (_AB * d / AB)) / H
	Trans[1][1] = _H / H
	Trans[1][2] = (h_3 - (h_2 * d / AB)) / H
	
	Trans[2][0] = mat.A(Trans, 2, 0)
	Trans[2][1] = mat.A(Trans, 2, 1)
	Trans[2][2] = mat.A(Trans, 2, 2)
	temp := math.Sqrt(math.Pow(Trans[2][0], 2) + math.Pow(Trans[2][1], 2) + math.Pow(Trans[2][2], 2))
	Trans[2][0] /= temp
	Trans[2][1] /= temp
	Trans[2][2] /= temp

	// New method
	p := (AB + AC + BC) / 2
	S := math.Sqrt(p * (p - AB) * (p - AC) * (p - BC))
	alpha := p1
	bravo := p2
	charlie := p3
	delta := BC
	echo := AC
	fox := AB
	Delta := bravo*bravo+charlie*charlie-delta*delta
	Echo := alpha*alpha+charlie*charlie-echo*echo
	Fox :=alpha*alpha+bravo*bravo-fox*fox
	V := math.Sqrt(4*alpha*alpha*bravo*bravo*charlie*charlie - alpha*alpha*Delta*Delta - bravo*bravo*Echo*Echo - charlie*charlie*Fox*Fox + Delta*Echo*Fox)/12
	Height := V * 3 / S
	d1 := math.Sqrt(p1 * p1 - Height * Height)
	d2 := math.Sqrt(p2 * p2 - Height * Height)
	d3 := math.Sqrt(p3 * p3 - Height * Height)
	COS_Alpha := (d1 * d1 + AB * AB - d2 * d2) / (2 * d1 * AB)
	Y := 0.0
	X := d1 * COS_Alpha
	if (1 - COS_Alpha) < 0.000000001 {
		Y = 0.0
	} else {
		Y = math.Sqrt(d1 * d1 - X * X)
	}
	
	d4 := math.Pow((d - X), 2) + math.Pow((H - Y), 2)
	d5 := math.Pow((d - X), 2) + math.Pow((H + Y), 2)
	e1 := math.Abs(math.Sqrt(d4) - d3)
	e2 := math.Abs(math.Sqrt(d5) - d3)
	if e2 < e1 {
		log.Println("below")
		Y = -Y
	}

	
	XYZ := mat.Zeros(1,3)
	XYZ[0][0] = X
	XYZ[0][1] = Y
	XYZ[0][2] = Height
	
	POS = mat.Mat_mult(XYZ, Trans)
	if mode == 0 {
		POS = Xyz2pos(POS)
	}
	
	return
}

func Xyz2pos (A mat.Matrix) mat.Matrix {
	POS := mat.Mat_mult(A, _Trans)
	POS[0][0] += Satellites_pos[0][0]
	POS[0][1] += Satellites_pos[0][1]
	POS[0][2] += Satellites_pos[0][2]
	return POS
}

func Pos_solu (Distance, _Satellites_pos mat.Matrix) mat.Matrix {
	Satellites_pos = _Satellites_pos
	guess, pl_xyz := Get_probable_pos(Distance, _Satellites_pos, 1)
	alpha := mat.Zeros(NUMOFPL, NUMOFPL)
	for i := 0; i < NUMOFPL; i++ {
		alpha[i][0] = 0.0;
		alpha[i][1] = 0.0;
		alpha[i][2] = 0.0;
	}
	err, count, rho, drho := 10.0, 0, mat.Zeros(1, 3), mat.Zeros(3, 1)
	
	for ; err > 0.1; count++ {
		for i := 0; i < NUMOFPL; i++ {
			rho[0][i] = math.Sqrt(math.Pow((guess[0][0] - pl_xyz[i][0]), 2) + math.Pow((guess[0][1] - pl_xyz[i][1]), 2) + math.Pow((guess[0][2] - pl_xyz[i][2]), 2));
		}
		for i := 0; i < NUMOFPL; i++ {
			alpha[i][0] = (guess[0][0] - pl_xyz[i][0]) / rho[0][i];
			alpha[i][1] = (guess[0][1] - pl_xyz[i][1]) / rho[0][i];
			alpha[i][2] = (guess[0][2] - pl_xyz[i][2]) / rho[0][i];
		}
		for i := 0; i < NUMOFPL; i++ {
			drho[i][0] = Distance[0][i] - rho[0][i];
		}
		// Be careful with drho
		dX := mat.Mat_mult(mat.Pinv(alpha), drho)
		guess[0][0] += dX[0][0]
		guess[0][1] += dX[1][0]
		guess[0][2] += dX[2][0]
		err = math.Sqrt(math.Pow(dX[0][0], 2) + math.Pow(dX[1][0], 2) + math.Pow(dX[2][0], 2))
		if err >= math.Pow(10,10) {
			log.Fatal("Failed to Calculate")
		}
	}
	return Xyz2pos(guess)
}