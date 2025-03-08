package q2234

import "testing"

func Test_maximumBeauty(t *testing.T) {
	type args struct {
		flowers    []int
		newFlowers int64
		target     int
		full       int
		partial    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test1",
			args: args{
				flowers:    []int{1, 3, 1, 1},
				newFlowers: 7,
				target:     6,
				full:       12,
				partial:    1,
			},
			want: 14,
		},
		{
			name: "test2",
			args: args{
				flowers:    []int{2, 4, 5, 3},
				newFlowers: 10,
				target:     5,
				full:       2,
				partial:    6,
			},
			want: 30,
		},
		{
			name: "test3",
			args: args{
				flowers: []int{
					22323, 64818, 97718, 14354, 27837, 6347, 43299, 23010, 18590, 12706,
					1579, 52401, 23473, 82978, 1012, 2248, 50247, 755, 12672, 57870, 90646,
					87848, 71069, 82723, 83385, 66909, 39266, 97768, 62453, 30454, 68978,
					88590, 11213, 74010, 65683, 75460, 3118, 98281, 28128, 84992, 52206, 92770,
					74240, 33266, 41603, 19267, 36991, 86658, 43834, 84533, 75187, 31502, 61181,
					31333, 37324, 13352, 51735, 89812, 56745, 44204, 85482, 70358, 48830, 91989,
					82778, 34042, 3293, 63626, 41301, 43763, 39681, 91917, 40165, 57944, 34357,
					22395, 26084, 21666, 40781, 28998, 12385, 10720, 66853, 42324, 28327, 30125,
					15522, 12223},
				newFlowers: 997843,
				target:     100000,
				full:       4880,
				partial:    45790,
			},
			want: 2080606020,
		},
		{
			name: "test4",
			args: args{
				flowers:    []int{20, 1, 15, 17, 10, 2, 4, 16, 15, 11},
				newFlowers: 2,
				target:     20,
				full:       10,
				partial:    2,
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumBeauty(tt.args.flowers, tt.args.newFlowers, tt.args.target, tt.args.full, tt.args.partial); got != tt.want {
				t.Errorf("maximumBeauty() = %v, want %v", got, tt.want)
			}
		})
	}
}
