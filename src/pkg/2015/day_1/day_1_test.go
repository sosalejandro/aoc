package day_1

import "testing"

func TestFloor_ReadFloorInstructions(t *testing.T) {
	type args struct {
		instructions string
	}
	tests := []struct {
		name string
		f    *Floor
		args args
		want Floor
	}{
		{
			name: "Test 1",
			f:    CreateFloor(),
			args: args{
				instructions: "(())",
			},
			want: 0,
		},
		{
			name: "Test 2",
			f:    CreateFloor(),
			args: args{
				instructions: "()()",
			},
			want: 0,
		},
		{
			name: "Test 3",
			f:    CreateFloor(),
			args: args{
				instructions: "(((",
			},
			want: 3,
		},
		{
			name: "Test 3",
			f:    CreateFloor(),
			args: args{
				instructions: "(()(()(",
			},
			want: 3,
		},
		{
			name: "Test 3",
			f:    CreateFloor(),
			args: args{
				instructions: "))(((((",
			},
			want: 3,
		},
		{
			name: "Test 3",
			f:    CreateFloor(),
			args: args{
				instructions: "())",
			},
			want: -1,
		},
		{
			name: "Test 3",
			f:    CreateFloor(),
			args: args{
				instructions: "))(",
			},
			want: -1,
		},
		{
			name: "Test 3",
			f:    CreateFloor(),
			args: args{
				instructions: ")))",
			},
			want: -3,
		},
		{
			name: "Test 3",
			f:    CreateFloor(),
			args: args{
				instructions: ")())())",
			},
			want: -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.f.ReadFloorInstructions1(tt.args.instructions)
			if *tt.f != tt.want {
				t.Errorf("Floor.ReadFloorInstructions() = %v, want %v", *tt.f, tt.want)
			}
		})
	}
}

func TestFloor_ReadFloorInstructions2(t *testing.T) {
	type args struct {
		instructions string
	}
	tests := []struct {
		name string
		f    *Floor
		args args
		want int
	}{
		{
			name: "Test 1",
			f:    CreateFloor(),
			args: args{
				instructions: ")",
			},
			want: 1,
		},
		{
			name: "Test 2",
			f:    CreateFloor(),
			args: args{
				instructions: "()())",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.ReadFloorInstructions2(tt.args.instructions); got != tt.want {
				t.Errorf("ReadFloorInstructions2() = %v, want %v", got, tt.want)
			}
		})
	}
}
