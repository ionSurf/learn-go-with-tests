package dictionary

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrWordExists      = DictionaryErr("word already exists in dictionary")
	ErrNotFound        = DictionaryErr("could not find the word you were looking for")
	ErrWordDoesNotExis = DictionaryErr("could not update value because word does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(query string) (result string, err error) {
	definition, ok := d[query]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) (err error) {
	_, searchErr := d.Search(word)

	switch searchErr {
	case ErrNotFound:
		d[word] = definition
		err = nil
	case nil:
		err = ErrWordExists
		return
	default:
		err = searchErr
	}

	return
}

func (d Dictionary) Update(word, definition string) (err error) {
	_, err = d.Search(word)

	if err == nil {
		d[word] = definition
	} else {
		err = ErrWordDoesNotExis
	}

	return
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
