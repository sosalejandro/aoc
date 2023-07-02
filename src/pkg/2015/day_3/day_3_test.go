package day_3

import (
	"testing"
)

func TestGrid_ParseGrid(t *testing.T) {
	type fields struct {
		Houses map[int]*House
	}
	type args struct {
		input string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Parse grid test 1",
			fields: fields{
				Houses: map[int]*House{},
			},
			args: args{
				input: ">",
			},
			want: 2,
		},
		{
			name: "Parse grid test 2",
			fields: fields{
				Houses: map[int]*House{},
			},
			args: args{
				input: "^>v<",
			},
			want: 4,
		},
		{
			name: "Parse grid test 3",
			fields: fields{
				Houses: map[int]*House{},
			},
			args: args{
				input: "^v^v^v^v^v",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := CreateGrid()
			g.ParseGrid(tt.args.input)
			result := len(g.Houses)
			if result != tt.want {
				t.Errorf("ParseGrid() = %v, want %v houses", result, tt.want)
			}
		})
	}
}

func TestGrid_ParseGrid2(t *testing.T) {
	type fields struct {
		Houses map[string]*House
	}
	type args struct {
		input string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Parse grid test 1",
			fields: fields{
				Houses: map[string]*House{},
			},
			args: args{
				input: "^v",
			},
			want: 3,
		},
		{
			name: "Parse grid test 2",
			fields: fields{
				Houses: map[string]*House{},
			},
			args: args{
				input: "^>v<",
			},
			want: 3,
		},
		{
			name: "Parse grid test 3",
			fields: fields{
				Houses: map[string]*House{},
			},
			args: args{
				input: "^v^v^v^v^v",
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := CreateGrid()
			g.ParseGrid2(tt.args.input)
			result := len(g.Houses)
			if result != tt.want {
				t.Errorf("ParseGrid() = %v, want %v houses", result, tt.want)
			}
		})
	}
}
