package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("known word", func(t *testing.T) {
		result, _ := dictionary.Search("test")
		expected := "this is a test"

		assertStrings(t, result, expected)

	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("Expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "this is another test"
		dictionary.Add(key, value)

		assertDefinition(t, dictionary, key, value)
	})
	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is another test"
		dictionary := Dictionary{key: value}

		err := dictionary.Add(key, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is another test"
		dictionary := Dictionary{key: value}

		err := dictionary.Update(key, "new test")

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, "new test")
	})
	t.Run("new word", func(t *testing.T) {
		key := "test"
		value := "this is a test"
		dictionary := Dictionary{}

		err := dictionary.Update(key, value)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	dictionary := Dictionary{key: "a value"}

	dictionary.Delete(key)

	_, err := dictionary.Search(key)
	assertError(t, err, ErrNotFound)
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, value string) {
	t.Helper()

	result, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, result, value)

}

func assertStrings(t testing.TB, result, expected string) {
	t.Helper()

	if result != expected {
		t.Errorf("got: %q, but expected %q, given %q", result, expected, "test")
	}
}

func assertError(t testing.TB, result, expected error) {
	t.Helper()

	if result != expected {
		t.Errorf("got: %q, but expected %q, given %q", result, expected, "test")
	}
}
