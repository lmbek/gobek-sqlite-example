package png

import (
	"errors"
	"net/http"
	"os"
)

func Get(response http.ResponseWriter, request *http.Request) (any, error) {
	file, err := os.Open("./data/test.png")
	if err != nil {
		// Handle the error if file cannot be opened
		return "", errors.New("Could not open ./data/test.png")
	}
	defer file.Close()

	// Get file information
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return nil, errors.New("Could not get file information")
	}
	response.Header().Set("Content-Type", "image/png")
	http.ServeContent(response, request, "test.png", fileInfo.ModTime(), file)
	return nil, errors.New("not json response")
}
