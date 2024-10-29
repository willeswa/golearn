package main

import (
	"fmt"
	"testing"
)

func TestInteger(t *testing.T){
	t.Run("call main function", func(t *testing.T) {
		main()
	})

	t.Run("add two numbers", func(t *testing.T){
		sum := Add(2, 2)
		expected := 4

		if(sum != expected){
			t.Errorf("expected '%d' but got '%d' ", expected, sum)
		}
	})
}

func ExampleAdd(){
	sum := Add(1, 6)
	fmt.Println(sum)
	// Output: 7
}