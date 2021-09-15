package slices

import (
	"reflect"
	"testing"
)

func Test_orderSlice(t *testing.T) {
	type args struct {
		values []int
		order  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "must order values ...",
			args: args{
				values: []int{10, 5, 12, 20, 25},
				order:  []int{0, 1, 2, 1, 2},
			},
			want: []int{10, 20, 25, 5, 12},
		},
		{
			name: "must order values 2...",
			args: args{
				values: []int{10, 5, 12, 20, 25},
				order:  []int{0, 1, 2, 3, 4},
			},
			want: []int{10, 5, 12, 20, 25},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orderSlice(tt.args.values, tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("orderSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
