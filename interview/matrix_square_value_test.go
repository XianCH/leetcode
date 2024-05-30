package interview

import "testing"

func TestMatrixSquareXvalue(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   int
	}{
		{
			matrix: [][]int{
				{6, 1, 1},
				{1, 6, 1},
				{1, 1, 6},
			},
			want: 10,
		},
		{
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: 0,
		},
		{
			matrix: [][]int{
				{3, 2, 1},
				{1, 2, 3},
				{3, 1, 2},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		got := matrixSquareXvalue(tt.matrix)
		if got != tt.want {
			t.Errorf("matrixSquareXvalue(%v)=%d;want %d", tt.matrix, got, tt.want)
		}
	}
}
