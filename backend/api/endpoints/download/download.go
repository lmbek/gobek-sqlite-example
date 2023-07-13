package download

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func Start(response http.ResponseWriter, request *http.Request) (any, error) {
	fmt.Printf("%v started download of %v \n", request.RemoteAddr, "./data/test.txt")
	file, err := os.Open("./data/test.txt")
	if err != nil {
		// Handle the error if file cannot be opened
		return "", errors.New("Could not open ./data/test.txt")
	}
	defer file.Close()

	// Get file information
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return nil, errors.New("Could not get file information")
	}
	response.Header().Set("Content-Type", "application/octet-stream")
	response.Header().Set("Content-Disposition", "attachment; filename=test.txt")
	http.ServeContent(response, request, "test.txt", fileInfo.ModTime(), file)

	return nil, errors.New("not json response")
}
