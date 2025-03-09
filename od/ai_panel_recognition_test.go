package od

import "testing"

func TestAiPanelRecognition(t *testing.T) {
	type args struct {
		lights [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				lights: [][]int{
					{1, 0, 0, 2, 2},
					{2, 6, 1, 8, 3},
					{3, 3, 2, 5, 4},
					{5, 5, 4, 7, 6},
					{4, 0, 4, 2, 6},
				},
			},
			want: "1 2 3 4 5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AiPanelRecognition(tt.args.lights); got != tt.want {
				t.Errorf("AiPanelRecognition() = %v, want %v", got, tt.want)
			}
		})
	}
}
