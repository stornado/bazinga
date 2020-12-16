package compare

import "testing"

func TestMin(t *testing.T) {
	type args struct {
		a      int
		b      int
		others []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"a=b", args{3, 3, nil}, 3},
		{"a<b", args{3, 4, nil}, 3},
		{"a>b", args{4, 3, nil}, 3},
		{"a<b<c", args{3, 4, []int{5}}, 3},
		{"a<b<c>d", args{3, 4, []int{5, 2}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.a, tt.args.b, tt.args.others...); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}
