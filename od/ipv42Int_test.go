package od

import "testing"

func TestIpv42Int(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				"128#0#255#255",
			},
			want: 2147549183,
		},
		{
			name: "t2",
			args: args{
				"1#0#0#0",
			},
			want: 16777216,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ipv42Int(tt.args.ip); got != tt.want {
				t.Errorf("Ipv42Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
