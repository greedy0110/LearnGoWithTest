package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dic := Dictionary{}
		word := "test"
		desc := "this is just a test"
		dic.Add(word, desc)

		assertDefinition(t, dic, word, desc)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		desc := "this is just a test"
		dic := Dictionary{word: desc}
		err := dic.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dic, word, desc)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		desc := "this is just a test"
		dic := Dictionary{word: desc}
		newDesc := "new test"
		err := dic.Update(word, newDesc)

		assertError(t, err, nil)
		assertDefinition(t, dic, word, newDesc)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		dic := Dictionary{}
		newDesc := "new test"
		err := dic.Update(word, newDesc)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dic := Dictionary{word: "test"}
	dic.Delete(word)

	_, err := dic.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertDefinition(t testing.TB, d Dictionary, word, desc string) {
	got, err := d.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != desc {
		t.Errorf("got %q want %q", got, desc)
	}
}
