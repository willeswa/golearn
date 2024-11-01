package main

import "testing"

var defition =  "A beautiful name given to an awesome girl"

func TestDictionary(t *testing.T){

	assertStrings := func (t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("want %s got %s", want, got)
		}
	}

	t.Run("test custom type", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a custom type"}
		got, _ := dictionary.Search("test")
		want :=  "this is a custom type"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{}
		_, err := dictionary.Search("test")
		
		if err == nil {
			t.Fatal("expected an error")
		}
		assertStrings(t, err.Error(), ErrorNotFound.Error())
	})

	t.Run("add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("Natalie", defition)
		got, _ := dictionary.Search("Natalie")
		assertStrings(t, got, defition)
	})

	
}