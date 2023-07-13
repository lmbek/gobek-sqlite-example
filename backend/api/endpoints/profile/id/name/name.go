package name

import (
	"errors"
)

// TODO: make everything into endpoints, so even if we just want to get name, we will handle it as endpoint, just like identifier
func Get(id string) (any, error) {
	if id == "995" {
		return "LMB", nil
	} else {
		return nil, errors.New("user id does not exist")
	}
}
