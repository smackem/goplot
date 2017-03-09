package calc

import (
	"reflect"
	"testing"
)

func TestFunction_Eval(t *testing.T) {
	tests := []struct {
		name   string
		f      Function
		steps  int
		wantXs []Number
		wantYs []Number
	}{
		{
			name: "x[1..11]->x",
			f: Function{
				Param: "x",
				lower: numberExpr{number: 1},
				upper: numberExpr{number: 11},
				body:  identExpr{ident: "x"},
			},
			steps:  10,
			wantXs: []Number{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			wantYs: []Number{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "x[0..4]->x*2",
			f: Function{
				Param: "x",
				lower: numberExpr{number: 0},
				upper: numberExpr{number: 4},
				body:  timesExpr{left: identExpr{ident: "x"}, right: numberExpr{number: 2}},
			},
			steps:  4,
			wantXs: []Number{0, 1, 2, 3},
			wantYs: []Number{0, 2, 4, 6},
		},
		{
			name: "x[0..0]->x",
			f: Function{
				Param: "x",
				lower: numberExpr{number: 0},
				upper: numberExpr{number: 0},
				body:  identExpr{ident: "x"},
			},
			steps:  3,
			wantXs: []Number{0, 0, 0},
			wantYs: []Number{0, 0, 0},
		},
		{
			name: "x[0..1]->x",
			f: Function{
				Param: "x",
				lower: numberExpr{number: 0},
				upper: numberExpr{number: 1},
				body:  identExpr{ident: "x"},
			},
			steps:  2,
			wantXs: []Number{0, 0.5},
			wantYs: []Number{0, 0.5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotXs, gotYs := tt.f.Eval(tt.steps)
			if !reflect.DeepEqual(gotXs, tt.wantXs) {
				t.Errorf("Function.Eval() gotXs = %v, want %v", gotXs, tt.wantXs)
			}
			if !reflect.DeepEqual(gotYs, tt.wantYs) {
				t.Errorf("Function.Eval() gotYs = %v, want %v", gotYs, tt.wantYs)
			}
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		src     string
		want    *Function
		wantErr bool
	}{
		{
			name: "x[0..10]->x",
			src:  "x[0..10]->x",
			want: &Function{
				Param: "x",
				lower: numberExpr{number: 0},
				upper: numberExpr{number: 10},
				body:  identExpr{ident: "x"},
			},
			wantErr: false,
		},
		{
			name:    "x[0->x",
			src:     "x[0->x",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
