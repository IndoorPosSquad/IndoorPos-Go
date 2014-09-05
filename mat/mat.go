package mat

import (
	//"errors"
)

type Matrix [][]float64

func Empty(A Matrix) bool {
	if len(A) == 0 || len(A[0]) == 0 {
		return true
	} else {
		return false
	}
}

func Identity(A, B Matrix) bool {
	rows_A, cols_A := len(A), len(A[0])
	rows_B, cols_B := len(B), len(B[0])
	if rows_A != rows_B || cols_A != cols_B {
		return false
	} else {
		for i := 0; i < rows_A; i++ {
			for j :=0; j < cols_A; j++ {
				if A[i][j] != B[i][j] {
					return false
				}
			}
		}
	}
	return true
}

func copyMatrix(B Matrix) Matrix {
	rows, cols := len(B), len(B[0])
	A := Zeros(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			A[i][j] = B[i][j]
		}
	}
	return A
}

func CopyMatrix(B Matrix) Matrix {
	rows, cols := len(B), len(B[0])
	A := Zeros(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			A[i][j] = B[i][j]
		}
	}
	return A
}

func Inv(A Matrix) Matrix {
	//Cannot Calculate Adj with only 1 number
	if len(A) == 1 && len(A[0]) == 1 {
		C := Zeros(1, 1)
		C[0][0] = 1 / A[0][0]
		return C
	}
	return Div(Adj(A), Det(A))
}

func Pinv(A Matrix) Matrix {
	//X' * inv(X * X')
	return Mat_mult(T(A), Inv(Mat_mult(A, T(A))))
}

func Div(B Matrix, n float64) Matrix {
	rows, cols := len(B), len(B[0])
	C := Zeros(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			C[i][j] = B[i][j] / n
		}
	}
	return C
}

func T(B Matrix) (Matrix) {
	rows, cols := len(B), len(B[0])
	A := Zeros(cols, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			A[j][i] = B[i][j]
		}
	}
	return A
}

func Adj(B Matrix) Matrix {
	rows, cols := len(B), len(B[0])
	A := copyMatrix(B)
	//newcopy(A, B)
	C := Zeros(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++{
			//_ = Remove(A, i, j)
			//fmt.Println("C[" , i , "][" , j , "]=" , Det(Remove(A, i, j) ))
			if (1 & (i + j)) == 1 {
				C[i][j] = - Det(remove(A, i, j))
				//fmt.Println("C[" , i , "][" , j , "]=" , -Det(remove(A, i, j) ))
			} else {
				C[i][j] = Det(remove(A, i, j))
				//fmt.Println("C[",i,"][",j,"]=",Det(remove(A, i, j)))
			}
		}
	}
	return T(C)
}

func A(B Matrix, i, j int) float64 {
	A := copyMatrix(B)
	//_ = Remove(A, i, j)
	//fmt.Println("C[" , i , "][" , j , "]=" , Det(Remove(A, i, j) ))
	if (1 & (i + j)) == 1 {
		return - Det(remove(A, i, j))
		//fmt.Println("C[" , i , "][" , j , "]=" , -Det(remove(A, i, j) ))
	} else {
		return Det(remove(A, i, j))
		//fmt.Println("C[",i,"][",j,"]=",Det(remove(A, i, j)))
	}
}

func remove(A Matrix, i, j int) Matrix {
	rows := len(A)
	//slice := make(Matrix, rows)
	slice := copyMatrix(A)
	//newcopy(slice, A)
	slice = append(slice[:i], slice[i+1:]...)
	rows -= 1
	for i := 0; i < rows; i++ {
		slice[i] = append(slice[i][:j], slice[i][j+1:]...)
	}
	return slice
}

func MakeMatrix(n int) Matrix {
	m := make(Matrix, n)
	for i := 0; i < n; i++ {
		m[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			m[i][j] = float64(i*n + j + 1)
		}
	}
	return m
}

func Mat_mult(A, B Matrix) (C Matrix) {
	k := len(B)
	if k != len(A[0]) {
		return Zeros(1, 1)//, errors.New("math: Matrix doesnt match")
	}
	i := len(A)
	//fmt.Println("len(A)",len(A))
	j := len(B[0])
	C = Zeros(i, j)
	matmult(nil, A, B, C, 0, i, 0, j, 0, k, 8)
	return C
}

func Zeros(x, y int) Matrix {
	m := make(Matrix, x)
	for i := 0; i < x; i++ {
		m[i] = make([]float64, y)
		for j := 0; j < y; j++ {
			m[i][j] = float64(0)
		}
	}
	return m
}

//k是A的列数orB的行数
//i是A的行数
//j是B的列数
func matmult(done chan<- struct{}, A, B, C Matrix, i0, i1, j0, j1, k0, k1, threshold int) {
	di := i1 - i0
	dj := j1 - j0
	dk := k1 - k0
	if di >= dj && di >= dk && di >= threshold {
		// divide in two by y axis
		mi := i0 + di/2
		done1 := make(chan struct{}, 1)
		go matmult(done1, A, B, C, i0, mi, j0, j1, k0, k1, threshold)
		matmult(nil, A, B, C, mi, i1, j0, j1, k0, k1, threshold)
		<-done1
	} else if dj >= dk && dj >= threshold {
		// divide in two by x axis
		mj := j0 + dj/2
		done1 := make(chan struct{}, 1)
		go matmult(done1, A, B, C, i0, i1, j0, mj, k0, k1, threshold)
		matmult(nil, A, B, C, i0, i1, mj, j1, k0, k1, threshold)
		<-done1
	} else if dk >= threshold {
		// divide in two by "k" axis
		// deliberately not parallel because of data races
		mk := k0 + dk/2
		matmult(nil, A, B, C, i0, i1, j0, j1, k0, mk, threshold)
		matmult(nil, A, B, C, i0, i1, j0, j1, mk, k1, threshold)
	} else {
		// the matrices are small enough, compute directly
		for i := i0; i < i1; i++ {
			for j := j0; j < j1; j++ {
				for k := k0; k < k1; k++ {
					C[i][j] += A[i][k] * B[k][j]
				}
			}
		}
	}
	if done != nil {
		done <- struct{}{}
	}
}

func Det(B Matrix) (float64) {

	n := len(B)
	
	if n != len(B[0]) {
		return 0
	}
	
	A := copyMatrix(B)
	//newcopy(A, B)
	sign, ret, j := 0, 1.0, 0
	
	for i := 0; i < n; i++ {
		if A[i][i] == 0.0 {
			for j = i+1; j < n; j++ {
				if A[j][i] != 0.0 {
					break
				}
			}
			if j == n {
				return 0
			}
			for k := i; k < n; k++ {
				t := A[i][k]
				A[i][k] = A[j][k]
				A[j][k] = t
			}
			sign++;
		}
		ret *= A[i][i]
		for k := i+1; k < n; k++ {
			A[i][k] /= A[i][i]
		}
		for j := i+1; j < n; j++ {
			for k := i+1; k < n; k++ {
				A[j][k] -= A[j][i] * A[i][k]
			}
		}
	}
	if (sign & 1) == 1 {
		ret = -ret
	}
	return ret
}