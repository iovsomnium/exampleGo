package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errWordExists = errors.New("That word already exists")
	errCantUpdate = errors.New("Can't update non-existing word")
)

//Search for word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

//Add new word to the dict
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil
}

// Update the word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	if err == nil {
		d[word] = def
	} else if err == errNotFound {
		return errCantUpdate
	}
	return nil
}
