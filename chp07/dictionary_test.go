package dictionary

import "testing"

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

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func assertError(t testing.TB, got, want error) {
	if got == nil {
		t.Fatal("Expecting an error")
	}
	t.Helper()
	if got != want {
		t.Errorf("Got error %q want error %q", got, want)
	}
}
