package basic

import "errors"

func Error() error {
	return errors.New("THIS ERROR")
}
