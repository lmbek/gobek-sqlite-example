package age

import (
	"errors"
)

func Get(id string) (any, error) {
	if id == "995" {
		return 28, nil
	} else {
		return nil, errors.New("user id does not exist")
	}
}
