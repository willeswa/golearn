package main

import "testing"

func TestHello(t *testing.T){
	got := Hello("Natalie")
	want := "Hello, Natalie!"

	if(got != want){
		t.Errorf("got %q want %q", got, want)
	}
}

func TestMain(m *testing.T){
	main()	
}