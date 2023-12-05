package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "This is just a test",
	}
	t.Run("word is in dictionary", func(t *testing.T) {
		query := "test"

		got, _ := dictionary.Search(query)
		want := "This is just a test"

		assertStrings(t, got, want)
	})
	t.Run("Word is not in dictionary", func(t *testing.T) {
		query := "NotInDictionary"

		_, err := dictionary.Search(query)

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Adds new word", func(t *testing.T) {
		word := "test"
		definition := "This is just a test"
		dictionary := Dictionary{}
		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("Adds existing word", func(t *testing.T) {
		word := "test"
		definition := "This is just a test"
		dictionary := Dictionary{
			word: definition,
		}
		err := dictionary.Add(word, definition)
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update existing word", func(t *testing.T) {
		word := "test"
		definition := "New word definition"
		dictionary := Dictionary{
			word: "Old word definition",
		}

		dictionary.Update(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("Update non-existing word", func(t *testing.T) {
		word := "test2"
		definition := "New word definition"
		dictionary := Dictionary{
			word: "Old word definition",
		}

		err := dictionary.Update("test", definition)

		assertError(t, err, ErrWordDoesNotExis)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{
		word: "dome definition",
	}
	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	assertError(t, err, ErrNotFound)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("Should have not returned an error")
	}
	assertStrings(t, got, definition)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("Got error %q want error %q", got, want)
	}
}
