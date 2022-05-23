package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	commands := &[]string{"KthLargest", "add", "add", "add", "add", "add"}
	values := &[][]int{{3, 4, 5, 8, 2}, {3}, {5}, {10}, {9}, {4}}
	for idx := 0; idx < b.N; idx++ {
		Run(commands, values)
	}
}
func TestRun(t *testing.T) {
	type args struct {
		commands *[]string
		values   *[][]int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example1",
			args: args{commands: &[]string{"KthLargest", "add", "add", "add", "add", "add"},
				values: &[][]int{{3, 4, 5, 8, 2}, {3}, {5}, {10}, {9}, {4}},
			},
			want: []string{"null", "4", "5", "5", "8", "8"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(tt.args.commands, tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
