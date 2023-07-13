package api

import (
	"api/types"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Init() {
	http.HandleFunc("/api/", ServeAPIUseGZip)
}

func ServeAPIUseGZip(response http.ResponseWriter, request *http.Request) {
	if !strings.Contains(request.Header.Get("Accept-Encoding"), "gzip") {
		ServeAPI(response, request)
	} else {
		response.Header().Set("Content-Encoding", "gzip")
		gzipping := gzip.NewWriter(response)
		defer gzipping.Close()
		ServeAPI(Writer{Writer: gzipping, ResponseWriter: response}, request)
	}
}

// ServeAPI - serves the api request by handling the response in a certain way
func ServeAPI(response http.ResponseWriter, request *http.Request) {
	// Redirect if not trailing slash
	// if the url contains / in the end, then we will ignore it for the API listener (return same result no matter if there is / or not
	//urlPath := request.URL.Path
	//	if hasTrailingSlash(urlPath) == false {
	//addTrailingSlashRedirect(urlPath, response, request)
	//	}
	//response.Header().Add("Cache-Control", "max-age=31536000, immutable")

	// setting cache to no cache
	response.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")

	// set response header as json
	response.Header().Set("Content-Type", "application/json")

	// response
	var jsonResponse any = types.JSONMessageResponse{false, "Not set"}

	// API endpoints
	data, err := HandleEndpoint(request.URL.Path, response, request)
	if err != nil {
		if err.Error() == "not json response" {
			return
		}
		jsonResponse = types.JSONMessageResponse{Success: false, Message: err.Error()}
	} else {
		jsonResponse = types.JSONDataResponse{Success: true, Data: data}
	}

	// Marshal JSON
	jsonFormattedResponse := jsonMarshalled(jsonResponse)

	// Send Response
	sendResponse(response, jsonFormattedResponse)
}

func jsonMarshalled(jsonResponse any) string {
	jsonFormattedResponse, err := json.Marshal(jsonResponse)
	//jsonResponseString, err := json.MarshalIndent(jsonResponse, "", "\t")
	if err != nil {
		fmt.Println("Could not json.Marshal - error: " + err.Error())
	}
	return string(jsonFormattedResponse)
}

func sendResponse(response http.ResponseWriter, jsonFormattedResponse string) {
	printText, err := fmt.Fprintln(response, jsonFormattedResponse)
	if err != nil {
		// if there was and error when writing response
		fmt.Println("Could not fmt.Fprintln - error: " + err.Error())
		fmt.Println("formattedPrintNewLine: " + string(rune(printText)))
	}
}

func hasTrailingSlash(urlPath string) bool {
	if urlPath[len(urlPath)-1] == '/' {
		return true
	} else {
		return false
	}
}

func addTrailingSlashRedirect(urlPath string, response http.ResponseWriter, request *http.Request) {
	http.Redirect(response, request, urlPath+"/", 301) // redirect url to without trailing slash, permanent redirect
}

type Writer struct {
	io.Writer
	http.ResponseWriter
}

func (gzipResponse Writer) Write(b []byte) (int, error) {
	return gzipResponse.Writer.Write(b)
}
