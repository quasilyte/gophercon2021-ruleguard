package example

import "errors"

func f() error {
	return errors.New("") // want `\Qempty errors are hard to debug`
}
