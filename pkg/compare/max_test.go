package compare

import "testing"

func TestMax(t *testing.T) {
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
		{"a<b", args{3, 4, nil}, 4},
		{"a>b", args{4, 3, nil}, 4},
		{"a<b<c", args{3, 4, []int{5}}, 5},
		{"a<b<c>d", args{3, 4, []int{5, 2}}, 5},
		{"a>b<c>d", args{6, 4, []int{5, 2}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.a, tt.args.b, tt.args.others...); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
