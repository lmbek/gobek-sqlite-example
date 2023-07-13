package id

import (
	"errors"
)

func Get(identifier string) (string, error) {
	if identifier == "995" {
		return "995", nil
		// types.JSONDataResponse{true, "995"}
	} else {
		return "", errors.New("user id does not exist")
		// types.JSONMessageResponse{false, "user id does not exist"}
	}
}
