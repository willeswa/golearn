package main

import "testing"

func TestDs(t *testing.T){
	t.Run("test sum array", func(t *testing.T){
		numbers := []int{1, 2, 3, 4, 5, 6}
		got := Sum(numbers)
		want := 21

		if(got != want ){
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}