package main

import (
	"testing"
)

func TestStructs(t *testing.T){
	checkArea := func(t testing.TB, shape Shape, want float64){
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("want %g got %g. Area for %#v", want, got, shape)
		}
	}

	checkPerimeter := func (t testing.TB, shape Shape, want float64)  {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("want %.2f got %.2f. Perimeter for: %#v", want, got, shape)
		}
	}

	areaStructs := []struct {
		shape Shape
		want float64
	}{
		{Rectangle{10.0, 5.0}, 50.0},
		{Cicle{7}, 153.93804002589985},
	}

	perimeterStructs := []struct {
		shape Shape
		want float64
	} {
		{Rectangle{10.0, 5.0}, 30.0},
		{Cicle{7}, 43.982297150257104},
	}

	for _, tt := range areaStructs {
		checkArea(t, tt.shape, tt.want)
	}
	
	for _, tt := range perimeterStructs {
		checkPerimeter(t, tt.shape, tt.want)
	}

}