package maps

const (
	ErrNotFound         = DictionaryErr("could find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	res, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return res, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, ok := d[word]
	if ok {
		return ErrWordExists
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, newDefintion string) error {
	_, ok := d[word]
	if !ok {
		return ErrWordDoesNotExist
	}

	d[word] = newDefintion
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
