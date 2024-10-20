package main

import "testing"

func TestHello(t *testing.T){
t.Run("Say hello to people", func(t *testing.T){
	got := Hello("Natalie", "")
	want := "Hello, Natalie"
	assertCorrectMessage(t, got, want)
})
t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T){
	got := Hello("", "")
	want := "Hello, World"
	assertCorrectMessage(t, got, want)
})

t.Run("Say hello in Spanish", func(t *testing.T){
	got := Hello("Natalie", "Spanish")
	want := "Hola, Natalie"
	assertCorrectMessage(t, got, want)
})

t.Run("Say hello in French", func(t *testing.T){
	got := Hello("Natalie", "French")
	want := "Bonjour, Natalie"
	assertCorrectMessage(t, got, want)
})
t.Run("Run main function", func(t *testing.T){
	main()
})
}

func assertCorrectMessage(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
