package main

import "testing"

func TestIteration(t *testing.T){
	t.Run("repeat character n times", func(t *testing.T){
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if(repeated != expected){
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

