package day_4

import "testing"

func TestFindAdventCoin(t *testing.T) {
	type args struct {
		input        string
		isAdventCoin func(Hash) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "IsAdventCoin: 'abcdef'",
			args: args{
				input:        "abcdef",
				isAdventCoin: Hash.IsAdventCoin,
			},
			want: 609043,
		},
		{
			name: "IsAdventCoin: 'pqrstuv'",
			args: args{
				input:        "pqrstuv",
				isAdventCoin: Hash.IsAdventCoin,
			},
			want: 1048970,
		},
		{
			name: "IsAdventCoin2: 'abcdef'",
			args: args{
				input:        "abcdef",
				isAdventCoin: Hash.IsAdventCoin2,
			},
			want: 6742839,
		},
		{
			name: "IsAdventCoin2: 'pqrstuv'",
			args: args{
				input:        "pqrstuv",
				isAdventCoin: Hash.IsAdventCoin2,
			},
			want: 5714438,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAdventCoin(tt.args.input, tt.args.isAdventCoin); got != tt.want {
				t.Errorf("FindAdventCoin() = %v, want %v", got, tt.want)
			}
		})
	}
}
