package tests

import (
	"api"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkAPIPath_Links(benchmark *testing.B) {
	// endpoint
	endpoint := "/api/links"

	// Reset the timer before running the code
	benchmark.ResetTimer()

	// create a mock response writer & request
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", endpoint, nil)

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint(endpoint, response, request)
	}
}

func BenchmarkAPIPath_User_NoIdentifier(benchmark *testing.B) {
	// endpoint
	endpoint := "/api/profile"

	// Reset the timer before running the code
	benchmark.ResetTimer()

	// create a mock response writer & request
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", endpoint, nil)

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint(endpoint, response, request)
	}
}

func BenchmarkAPIPath_User_Empty(benchmark *testing.B) {
	// Reset the timer before running the code
	benchmark.ResetTimer()

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint("/api/profile/", nil, nil)
	}
}

func BenchmarkAPIPath_User_995(benchmark *testing.B) {
	// endpoint
	endpoint := "/api/user/995"

	// Reset the timer before running the code
	benchmark.ResetTimer()

	// create a mock response writer & request
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", endpoint, nil)

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint(endpoint, response, request)
	}
}

func BenchmarkAPIPath_User_995_slash(benchmark *testing.B) {
	// endpoint
	endpoint := "/api/user/995/"

	// Reset the timer before running the code
	benchmark.ResetTimer()

	// create a mock response writer & request
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", endpoint, nil)

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint(endpoint, response, request)
	}
}

func BenchmarkAPIPath_User_995_name(benchmark *testing.B) {
	// endpoint
	endpoint := "/api/user/995/name"

	// Reset the timer before running the code
	benchmark.ResetTimer()

	// create a mock response writer & request
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", endpoint, nil)

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint(endpoint, response, request)
	}
}

func BenchmarkAPIPath_User_995_name_something(benchmark *testing.B) {
	// endpoint
	endpoint := "/api/user/995/name/something"

	// Reset the timer before running the code
	benchmark.ResetTimer()

	// create a mock response writer & request
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", endpoint, nil)

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint(endpoint, response, request)
	}
}

func BenchmarkAPIPath_User_NonExisting(benchmark *testing.B) {
	// endpoint
	endpoint := "/api/user/996"

	// Reset the timer before running the code
	benchmark.ResetTimer()

	// create a mock response writer & request
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", endpoint, nil)

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint(endpoint, response, request)
	}
}

func BenchmarkAPIPath_User_invalid_endpoint(benchmark *testing.B) {
	// endpoint
	endpoint := "/api/user/995/age/something/something/something/something"

	// Reset the timer before running the code
	benchmark.ResetTimer()

	// create a mock response writer & request
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", endpoint, nil)

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint(endpoint, response, request)
	}
}

func BenchmarkAPIPath_User_NonExisting_invalid_endpoint(benchmark *testing.B) {
	// endpoint
	endpoint := "/api/user/996/age/something/something/something/something"

	// Reset the timer before running the code
	benchmark.ResetTimer()

	// create a mock response writer & request
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", endpoint, nil)

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint(endpoint, response, request)
	}
}

func BenchmarkAPIPath_invalid_endpoint(benchmark *testing.B) {
	// endpoint
	endpoint := "/api/something/something/something/something"

	// Reset the timer before running the code
	benchmark.ResetTimer()

	// create a mock response writer & request
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", endpoint, nil)

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint(endpoint, response, request)
	}
}

func BenchmarkAPIPath_Users(benchmark *testing.B) {
	// endpoint
	endpoint := "/api/profiles"

	// Reset the timer before running the code
	benchmark.ResetTimer()

	// create a mock response writer & request
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", endpoint, nil)

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint(endpoint, response, request)
	}
}

func BenchmarkAPIPath_Places(benchmark *testing.B) {
	// endpoint
	endpoint := "/api/places"

	// Reset the timer before running the code
	benchmark.ResetTimer()

	// create a mock response writer & request
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", endpoint, nil)

	// Run the code b.N times
	for i := 0; i < benchmark.N; i++ {
		api.HandleEndpoint(endpoint, response, request)
	}
}
