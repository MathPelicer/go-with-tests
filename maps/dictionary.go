package maps

type Dictionary map[string]string

var (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update a word that does not exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (dictionary Dictionary) Search(word string) (string, error) {
	definition, isFound := dictionary[word]
	if !isFound {
		return "", ErrNotFound
	}

	return definition, nil
}

func (dictionary Dictionary) Add(key, value string) error {
	_, err := dictionary.Search(key)

	switch err {
	case ErrNotFound:
		dictionary[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (dictionary Dictionary) Update(key, value string) error {
	_, err := dictionary.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		dictionary[key] = value
	default:
		return err
	}

	return nil

}

func (dictionary Dictionary) Delete(key string) {
	delete(dictionary, key)
}
