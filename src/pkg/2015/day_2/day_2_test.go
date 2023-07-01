package day_2

import "testing"

func TestPresentBox_GetSmallestSideArea(t *testing.T) {
	type fields struct {
		Length int
		Width  int
		Height int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "2x3x4",
			fields: fields{Length: 2, Width: 3, Height: 4},
			want:   6,
		},
		{
			name:   "1x1x10",
			fields: fields{Length: 1, Width: 1, Height: 10},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PresentBox{
				Length: tt.fields.Length,
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			if got := p.GetSmallestSideArea(); got != tt.want {
				t.Errorf("GetSmallestSideArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPresentBox_GetSurfaceArea(t *testing.T) {
	type fields struct {
		Length int
		Width  int
		Height int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "2x3x4",
			fields: fields{Length: 2, Width: 3, Height: 4},
			want:   52,
		},
		{
			name:   "1x1x10",
			fields: fields{Length: 1, Width: 1, Height: 10},
			want:   42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PresentBox{
				Length: tt.fields.Length,
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			if got := p.GetSurfaceArea(); got != tt.want {
				t.Errorf("GetSurfaceArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPresentBox_GetWrappingPaperArea(t *testing.T) {
	type fields struct {
		Length int
		Width  int
		Height int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "2x3x4",
			fields: fields{Length: 2, Width: 3, Height: 4},
			want:   58,
		},
		{
			name:   "1x1x10",
			fields: fields{Length: 1, Width: 1, Height: 10},
			want:   43,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PresentBox{
				Length: tt.fields.Length,
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			if got := p.GetWrappingPaperArea(); got != tt.want {
				t.Errorf("GetWrappingPaperArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPresentBox_GetSmallestPerimeter(t *testing.T) {
	type fields struct {
		Length int
		Width  int
		Height int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "2x3x4",
			fields: fields{Length: 2, Width: 3, Height: 4},
			want:   10,
		},
		{
			name:   "1x1x10",
			fields: fields{Length: 1, Width: 1, Height: 10},
			want:   4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PresentBox{
				Length: tt.fields.Length,
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			if got := p.GetSmallestPerimeter(); got != tt.want {
				t.Errorf("GetSmallestPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPresentBox_GetVolume(t *testing.T) {
	type fields struct {
		Length int
		Width  int
		Height int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "2x3x4",
			fields: fields{Length: 2, Width: 3, Height: 4},
			want:   24,
		},
		{
			name:   "1x1x10",
			fields: fields{Length: 1, Width: 1, Height: 10},
			want:   10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PresentBox{
				Length: tt.fields.Length,
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			if got := p.GetVolume(); got != tt.want {
				t.Errorf("GetVolume() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRibbonCalculator_GetRibbonLength(t *testing.T) {
	type args struct {
		box *PresentBox
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2x3x4",
			args: args{box: &PresentBox{Length: 2, Width: 3, Height: 4}},
			want: 34,
		},
		{
			name: "1x1x10",
			args: args{box: &PresentBox{Length: 1, Width: 1, Height: 10}},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RibbonCalculator{}
			if got := r.GetRibbonLength(tt.args.box); got != tt.want {
				t.Errorf("GetRibbonLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
